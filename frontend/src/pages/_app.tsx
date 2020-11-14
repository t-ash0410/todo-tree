import React from 'react';
import { ThemeProvider, StylesProvider } from '@material-ui/styles';
import { CssBaseline, Container, makeStyles } from '@material-ui/core';
import theme from '../styles/theme';
import Header from '../components/organisms/header'

const useStyles = makeStyles((theme) => ({
    container: {
        paddingTop: theme.spacing(2),
    }
}));

const MyApp = ({ Component, pageProps }): JSX.Element => {
    const classes = useStyles();

    React.useEffect(() => {
        const jssStyles = document.querySelector('#jss-server-side')
        if (jssStyles && jssStyles.parentNode) {
            jssStyles.parentNode.removeChild(jssStyles)
        }
    }, []);

    return (
        <StylesProvider injectFirst>
            <ThemeProvider theme={theme}>
                <CssBaseline />
                <Header />
                <Container fixed className={classes.container}>
                    <Component {...pageProps} />
                </Container>
            </ThemeProvider>
        </StylesProvider>
    )
}

export default MyApp;
