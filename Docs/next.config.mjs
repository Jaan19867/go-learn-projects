import nextra from "nextra"

// Set up Nextra with its configuration
const withNextra = nextra({
  // ... Add Nextra-specific options here
//   theme: "nextra-theme-docs",
//   themeConfig: "./theme.config.mjs",
//   defaultShowCopyCode: true,
})

// Export the final Next.js config with Nextra included
export default withNextra({
  // ... Add regular Next.js options here
//   contentDirBasePath: "/docs",
})
