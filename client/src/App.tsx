import React from 'react'
import { hot } from 'react-hot-loader/root'
import CssBaseline from '@material-ui/core/CssBaseline'
import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'
import Header from './components/Header'
import Home from './containers/Home'
import { createStyles, Theme, withStyles, WithStyles } from '@material-ui/core/styles'

const styles = ({spacing}: Theme) => createStyles({
  main: {
    marginTop: 90,
    paddingLeft: spacing.unit * 5,
    paddingRight: spacing.unit * 5
  }
})

interface IProps extends WithStyles<typeof styles> {
}

const App = (props: IProps) => {
  return (
    <>
      <CssBaseline />

      <Grid container spacing={24}>
        <Grid item xs={12}>
          <Header />

          <main className={props.classes.main}>
            <Home />
          </main>
        </Grid>
      </Grid>
    </>
  )
}

export default hot(withStyles(styles)(App))
