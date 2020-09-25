import React, { Component } from 'react';
import './App.css';
import NavigationBar from './components/NavigationBar';
import Footer from './components/Footer';

// Object destructuring
// const StudentCard = ({name, age, major}) => {
//   return (
//     <div className="student">
//       <p className="student-name">{name}</p>
//       <p className="student-age">{age}</p>
//       <p className="student-major">{major}</p>
//     </div>
//   )
// }

class App extends Component {
  render(){
    return (
      <div className="App">
        <NavigationBar />
        {/* <StudentCard name="Irfan" age="24" major="International Relationship"/>       */}
        <Footer />
      </div>
    );
  }
}

export default App;
