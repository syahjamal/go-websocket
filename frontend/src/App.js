import React, { Component } from 'react';
import './App.css';
import NavigationBar from './components/NavigationBar';
import Footer from './components/Footer';
import Accordion from './components/Accordion';
import { connect, sendMsg } from './api';


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
    // instance of websocket connection as a class property
    // componentDidMount() {
    //   console.log("connecting")

    //     this.ws.onopen = () => {
    //     // on connecting, do nothing but log it to the console
    //     console.log('connected')
    //     }
  
    //     this.ws.onmessage = evt => {
    //     // listen to data sent from the websocket server
    //     const message = JSON.parse(evt.data)
    //     this.setState({dataFromServer: message})
    //     console.log("ini server")
    //     }
  
    //     this.ws.onclose = () => {
    //     console.log('disconnected')
    //     // automatically try to reconnect on connection loss
  
    //     }
  
    // }
  
    // render(){
    //   return(
    //     <ChildComponent websocket={this.ws} />
    //   )
    // }

  render(){
    return (
      <div className="App">
        <NavigationBar />
        {/* <Accordion /> */}
        <Footer />
      </div>
    );
  }
}

export default App;
