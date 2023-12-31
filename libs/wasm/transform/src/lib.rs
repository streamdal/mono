use protos::sp_steps_transform::TransformType;
use protos::sp_wsm::{WASMExitCode, WASMRequest};
use streamdal_wasm_transform::transform;

#[no_mangle]
pub extern "C" fn f(ptr: *mut u8, length: usize) -> u64 {
    // Read request
    let wasm_request = match common::read_request(ptr, length) {
        Ok(req) => req,
        Err(e) => panic!("unable to read request: {}", e), // TODO: Should write response here
    };

    // Validate request
    if let Err(err) = validate_wasm_request(&wasm_request) {
        panic!("invalid step: {}", err) // TODO: Should write response here
    }

    // Generate transform request
    let transform_request = generate_transform_request(&wasm_request);

    // Run request against transform
    let result = match wasm_request.step.transform().type_.unwrap() {
        TransformType::TRANSFORM_TYPE_REPLACE_VALUE => transform::overwrite(&transform_request),
        TransformType::TRANSFORM_TYPE_MASK_VALUE => transform::mask(&transform_request),
        TransformType::TRANSFORM_TYPE_OBFUSCATE_VALUE => transform::obfuscate(&transform_request),
        _ => {
            return common::write_response(
                None,
                None,
                WASMExitCode::WASM_EXIT_CODE_FAILURE,
                "Unknown transform type".to_string(),
            )
        }
    };

    // Inspect result and return potentially transformed payload
    match result {
        Ok(data) => common::write_response(
            Some(&data.into_bytes()),
            None,
            WASMExitCode::WASM_EXIT_CODE_SUCCESS,
            "Successfully transformed payload".to_string(),
        ),
        Err(err) => common::write_response(
            None,
            None,
            WASMExitCode::WASM_EXIT_CODE_FAILURE,
            format!("Unable to transform payload: {:?}", err),
        ),
    }
}

fn generate_transform_request(wasm_request: &WASMRequest) -> transform::Request {
    transform::Request {
        data: wasm_request.input_payload.clone(),
        path: wasm_request.step.transform().path.clone(),
        value: wasm_request.step.transform().value.clone(),
    }
}

fn validate_wasm_request(req: &WASMRequest) -> Result<(), String> {
    if req.input_payload.is_empty() {
        return Err("input cannot be empty".to_string());
    }

    if !req.step.has_transform() {
        return Err("transform is required".to_string());
    }

    let transform_type = req.step.transform().type_.enum_value_or_default();

    if transform_type == TransformType::TRANSFORM_TYPE_UNKNOWN {
        return Err("transform type cannot be unknown".to_string());
    }

    if req.step.transform().path.is_empty() {
        return Err("transform path cannot be empty".to_string());
    }

    Ok(())
}

/// # Safety
///
/// This is unsafe because it operates on raw memory; see `common/src/lib.rs`.
#[no_mangle]
pub unsafe extern "C" fn alloc(size: i32) -> *mut u8 {
    common::alloc(size)
}

/// # Safety
///
/// This is unsafe because it operates on raw memory; see `common/src/lib.rs`.
#[no_mangle]
pub unsafe extern "C" fn dealloc(pointer: *mut u8, size: i32) {
    common::dealloc(pointer, size)
}
