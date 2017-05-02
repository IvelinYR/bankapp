import React, {Component} from 'react';
import './App.css';
import Header from './components/Header/Header'

export default class App extends Component {
    render() {
        return (
            <div className="App">
                <div>
                    <Header/>
                    {this.props.children}
                </div>
            </div>
        );
    }
}