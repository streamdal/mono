import { DocSearchModal, useDocSearchKeyboardEvents } from "@docsearch/react";
import { createPortal } from "preact/compat";
import {
  MutableRef,
  useCallback,
  useEffect,
  useRef,
  useState,
} from "preact/hooks";

interface Props {
  lang?: string;
  labels: {
    modal: any;
    placeholder: string;
  };
}

export default function Search({ lang = "en", labels }: Props) {
  const [isOpen, setIsOpen] = useState(false);
  const searchButtonRef = useRef<any>(
    document.getElementById("docsearch-search-button")
  );
  const [initialQuery, setInitialQuery] = useState<string>();

  const onOpen = useCallback(() => {
    setIsOpen(true);
  }, [setIsOpen]);

  const onClose = useCallback(() => {
    setIsOpen(false);
  }, [setIsOpen]);

  const onInput = useCallback(
    (e: any) => {
      setIsOpen(true);
      setInitialQuery(e.key);
    },
    [setIsOpen, setInitialQuery]
  );

  useEffect(() => {
    searchButtonRef.current?.addEventListener("click", onOpen);
    return () => searchButtonRef.current?.removeEventListener("click", onOpen);
  }, [searchButtonRef.current, onOpen]);

  useDocSearchKeyboardEvents({
    isOpen,
    onOpen,
    onClose,
    onInput,
    searchButtonRef,
  });

  if (!isOpen) return null;

  return createPortal(
    <DocSearchModal
      initialQuery={initialQuery}
      initialScrollY={window.scrollY}
      onClose={onClose}
      appId={import.meta.env.PUBLIC_ALGOLIA_APP_ID}
      apiKey={import.meta.env.PUBLIC_ALGOLIA_API_KEY}
      indexName="streamdal"
      getMissingResultsUrl={({ query }) =>
        `https://github.com/withastro/docs/issues/new?title=Missing+results+for+query+%22${encodeURIComponent(
          query
        )}%22`
      }
      transformItems={(items) => {
        return items.map((item) => {
          // We transform the absolute URL into a relative URL to
          // work better on localhost, preview URLS.
          const a = document.createElement("a");
          a.href = item.url;
          const hash = a.hash === "#what-is-streamdal" ? "" : a.hash;
          return {
            ...item,
            url: `${a.pathname}${hash}`,
          };
        });
      }}
      placeholder={labels.placeholder}
    />,
    document.body
  );
}
