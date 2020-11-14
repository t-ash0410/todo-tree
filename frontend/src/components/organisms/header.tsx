import { Toolbar, IconButton, Typography, Box } from '@material-ui/core'
import MenuIcon from '@material-ui/icons/Menu';
import { makeStyles } from '@material-ui/core/styles';
import Link from 'next/link'

const Header = (): JSX.Element => {
  const styles = makeStyles((theme) => ({
    toolbar: {
      flexGrow: 1,
      background: '#333',
      color: theme.palette.primary.main
    },
    title: {
      flexGrow: 1,
      textDecoration: 'none',
      '&:hover': {
        cursor: 'pointer'
      }
    }
  }))();

  return (
    <Toolbar className={styles.toolbar}>
      <Box mr={1}>
        <IconButton edge="start" color="inherit" aria-label="menu">
          <MenuIcon />
        </IconButton>
      </Box>
      <Link href="/">
        <Typography variant="h6" className={styles.title}>
        ToDo
        </Typography>
      </Link>
    </Toolbar>
  )
}

export default Header;
