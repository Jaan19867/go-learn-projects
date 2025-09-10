import { useMDXComponents as getThemeComponents } from "nextra-theme-docs" // nextra-theme-blog or your custom theme

// Get the default MDX components
const themeComponents = getThemeComponents()

// Merge components
export function useMDXComponents(components) {
  return {
    ...themeComponents,
    h1: (props) => <h1 style={{ color: "red" }} {...props} />,
    Button: (props) => <button className="btn" {...props} />,
    ...components,
  }
}
