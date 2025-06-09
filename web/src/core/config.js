/**
 * 网站配置文件
 */
const greenText = (text) => `\x1b[32m${text}\x1b[0m`

const config = {
  appName: '电报账号',
  appLogo: 'logo.png',
  showViteLogo: true,
  logs: []
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
   greenText(
    console.log(
      `   ███████╗██╗  ██╗███████╗███╗   ███╗ ██████╗ ██╗  ██╗
   ██╔════╝╚██╗██╔╝██╔════╝████╗ ████║██╔═══██╗╚██╗██╔╝
   █████╗   ╚███╔╝ █████╗  ██╔████╔██║██║   ██║ ╚███╔╝ 
   ██╔══╝   ██╔██╗ ██╔══╝  ██║╚██╔╝██║██║   ██║ ██╔██╗ 
   ███████╗██╔╝ ██╗███████╗██║ ╚═╝ ██║╚██████╔╝██╔╝ ██╗
`,
    )
   )
  }
}

export default config
