import React from 'react'
import { withSiteData } from 'react-static'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'
import { Button } from '@material-ui/core'
import Grid from '@material-ui/core/Grid'
import Title from 'components/Title'
import BridgeList from 'components/BridgeList'
import ReactStaticLinkComponent from 'components/ReactStaticLinkComponent'
import matchRouteAndMapDispatchToProps from 'utils/matchRouteAndMapDispatchToProps'
import bridgesSelector from 'selectors/bridges'
import { fetchBridges } from 'actions'

export const Index = (props) => {
  const { bridges, bridgeCount, pageSize, bridgesError, fetchBridges, history, match } = props
  return <div>
    <Grid container spacing={8} alignItems='center'>
      <Grid item xs={9}>
        <Title>Bridges</Title>
      </Grid>
      <Grid item xs={3}>
        <Grid container justify='flex-end'>
          <Grid item>
            <Button
              variant='outlined'
              color='primary'
              component={ReactStaticLinkComponent}
              to={'/bridges/new'}
            >
              New Bridge
            </Button>
          </Grid>
        </Grid>
      </Grid>
    </Grid>
    <Grid container spacing={40}>
      <Grid item xs={12}>
        <BridgeList
          bridges={bridges}
          bridgeCount={bridgeCount}
          pageSize={pageSize}
          error={bridgesError}
          fetchBridges={fetchBridges}
          history={history}
          match={match}
        />
      </Grid>
    </Grid>
  </div>
}

Index.propTypes = {
  bridgeCount: PropTypes.number.isRequired,
  bridges: PropTypes.array.isRequired,
  bridgesError: PropTypes.string,
  pageSize: PropTypes.number
}

Index.defaultProps = {
  pageSize: 10
}

const mapStateToProps = state => {
  let bridgesError
  if (state.bridges.networkError) {
    bridgesError = 'There was an error fetching the bridges. Please reload the page.'
  }

  return {
    bridgeCount: state.bridges.count,
    bridges: bridgesSelector(state),
    bridgesError: bridgesError
  }
}

export const ConnectedIndex = connect(
  mapStateToProps,
  matchRouteAndMapDispatchToProps({fetchBridges})
)(Index)

export default withSiteData(ConnectedIndex)
