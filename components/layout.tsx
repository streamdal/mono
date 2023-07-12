import { Head } from "$fresh/runtime.ts";
import { NavBar } from "./nav/nav.tsx";
import { ComponentChildren } from "https://esm.sh/v128/preact@10.15.1/src/index.d.ts";

export type MetaProps = {
  title?: string;
  name?: string;
  description?: string;
};

export type LayoutProps = {
  children: ComponentChildren;
  meta?: MetaProps;
  hideNav?: boolean;
};
export const Layout = ({ children, hideNav, ...meta }: LayoutProps) => (
  <html lang="en">
    <Head>
      <title>{meta.title || "Snitch Console"}</title>
      <meta
        content={meta.description || "Streamdal's Snitch Console"}
        name="description"
      />
      <link rel="icon" type="image/png" href="/images/favicon.png" />
      <link
        rel="stylesheet"
        href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;600;700&family=Space+Grotesk:wght@400;500&display=swap"
      />
      <link
        rel="stylesheet"
        href="https://esm.sh/reactflow@11.7.4/dist/base.css"
      />
      <link
        rel="stylesheet"
        href="https://esm.sh/reactflow@11.7.4/dist/style.css"
      />
      <link rel="stylesheet" href="style.css" />
      <link
        href="https://cdnjs.cloudflare.com/ajax/libs/flowbite/1.7.0/flowbite.min.css"
        rel="stylesheet"
      />
    </Head>
    <body className="h-screen">
      <div className="flex flex-col w-full h-screen text-web bg-purple-50">
        {hideNav ? null : <NavBar />}
        <div className="mx-8 my-6">
          {children}
        </div>
      </div>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/flowbite/1.7.0/flowbite.min.js">
      </script>
    </body>
  </html>
);
