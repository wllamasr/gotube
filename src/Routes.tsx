import React from 'react'
import {BrowserRouter as Router, Route, Switch} from "react-router-dom";

export const Routes = () => {
    return <Router>
        <Switch>
            <Route path="/about">
            </Route>
            <Route path="/users">
            </Route>
            <Route path="/">
            </Route>
        </Switch>
    </Router>
}