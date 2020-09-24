import React, { Component } from 'react';
import './App.css';
import NavigationBar from './components/NavigationBar';
import Footer from './components/Footer';
import { connect } from './api';


class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      notifHistory: []
    }
  }
  
  componentDidMount() {
    connect((msg) => {
      console.log("New Message")
      this.setState(prevState => ({
        notifHistory: [...prevState.notifHistory, msg]
      }))
    })
  }

  render(){
    return (
      <div className="App">
        <NavigationBar />
        <Footer />
      </div>
    );
  }
}

export default App;
