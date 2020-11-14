import { createMuiTheme } from '@material-ui/core';
const theme = createMuiTheme({
  palette: {
    type: "dark",
    primary: {
      main: '#90caf9'
    },
    secondary: {
      main: '#f48fb1'
    },
    background: {
      default: '#222222'
    },
    contrastThreshold: 3,
    tonalOffset: 0.2
  }
});
export default theme;
