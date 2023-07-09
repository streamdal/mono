fn main() {
    println!("Overwrite example");
    overwrite();

    println!("Mask string example");
    mask_string();

    println!("Mask number example");
    mask_number();
}

fn overwrite() {
    let sample_json = r#"{"hello": "world"}"#;

    let mut req = protos::transform::TransformRequest::new();
    req.data = sample_json.into();
    req.path = "hello".to_string();
    req.value = r#""baz""#.to_string();

    let updated_json = snitch_transform::transform::overwrite(&req).unwrap();

    println!(
        "Input JSON: {} || Result JSON: {}",
        sample_json, updated_json,
    )
}

fn mask_string() {
    let sample_json = r#"{"hello": "world"}"#;

    let mut req = protos::transform::TransformRequest::new();
    req.path = "hello".to_string();
    req.data = sample_json.into();

    let updated_json = snitch_transform::transform::mask(&req).unwrap();

    println!(
        "Input JSON: {} || Result JSON: {}",
        sample_json, updated_json,
    )
}

fn mask_number() {
    let sample_json = r#"{"hello": 329328102938}"#;

    let mut req = protos::transform::TransformRequest::new();
    req.path = "hello".to_string();
    req.data = sample_json.into();

    let updated_json = snitch_transform::transform::mask(&req).unwrap();

    println!(
        "Input JSON: {} || Result JSON: {}",
        sample_json, updated_json,
    )
}
