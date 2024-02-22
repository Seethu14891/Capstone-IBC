
import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import { AuthRoute } from '../gaurds/auth-guard';
import { Login } from '../pages/Login';
import { Register } from '../pages/Register';
import { Dashboard } from '../pages/Dashboard';
import { NotFound } from '../pages/NotFound';
import { ViewProducts } from '../pages/ViewProducts';
import { Cart } from '../pages/Cart';

export const AppRoutes = {
  routes : [
    {
      path: "/",
      component: Login,
      exact: true
    },
    {
      path: "/register",
      component: Register,
      exact: true
    },
    {
      path: '/dashboard',
      component: Dashboard,
      auth: true
    },
    {
      path: '*',
      component: NotFound
    }
  ] 
}

export class RouteBody extends Component {
    constructor(props) {
      super(props);
      this.state = {
        body: this.prepareRoutes()
      };
    }
    prepareRoutes = () => {
      let  body = AppRoutes.routes.map((routesInfo, index) => {
        if(routesInfo.hasOwnProperty('auth') && routesInfo.auth === true) {
          return <AuthRoute key={index} path={routesInfo.path} component={routesInfo.component} />
        } else {
          return <Route key={index} path={routesInfo.path} component={routesInfo.component} exact />
        }
      });
      return body;
    }
    render() {
      return (
        <div> 
          <BrowserRouter>
            <Switch>
              { this.state.body }
            </Switch>
          </BrowserRouter>
        </div>
      )
    }
    
}

export const DashboardRoutes = {
  routes : [
    {
      path: "/dashboard/viewProducts",
      component: ViewProducts,
    },
    {
      path: "/dashboard/cart",
      component: Cart,
    }
  ] 
}

export class DashboardRoutesBody extends Component {
    constructor(props) {
      super(props);
      this.state = {
        body: this.prepareRoutes()
      };
    }
    prepareRoutes = () => {
      debugger;
      let  body = DashboardRoutes.routes.map((routesInfo, index) => {
        index = 'd'+ index;
        if(routesInfo.hasOwnProperty('auth') && routesInfo.auth === true) {
          return <AuthRoute key={index} path={routesInfo.path} component={routesInfo.component} />
        } else {
          return <Route key={index} path={routesInfo.path} component={routesInfo.component} exact />
        }
      });
      return body;
    }
    render() {
      console.log(this.state.body)
      return (
        <div> 
          { this.state.body }
        </div>
      )
    }
    
}

