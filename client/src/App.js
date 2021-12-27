import React, {useEffect, useState} from "react";
import {create, get} from "./http/membersAPI";
import "./App.css";

function App() {
  // const arr = [{"id":"1","name":"John", "email":"j.snow@company.com", "registration":"17.01.2019"},
  //   {"id":"2","name":"Dana", "email":"dana200@mail.com", "registration":"05.12.2020"},
  //   {"id":"3","name":"Henri", "email":"rouse@impress.com", "registration":"10.05.2021"},
  //   {"id":"4","name":"Brown", "email":"browm@gmail.com", "registration":"12.05.2021"}
  // ];

  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [members, setMembers] = useState([]);

  const add = async() => {
    await create(email,name)
  };
  const show = async() =>{
    const res = await get()
    setMembers(res.data)
  };
  const clear = () => {
    setName('');
    setEmail('')
  };


  return (
    <div className="App">
      <h1>Welcome to the club!</h1>
        <div className="add-form">
          <div className="fields">
            <div><label>Name: </label><input type="text" onChange={(e) => {setName(e.target.value)}} value={name}/> </div>
            <div><label>Email: </label><input type="text" onChange={(e) => {setEmail(e.target.value)}} value={email}/></div>
          </div>
          <div className="buttons">
            <button onClick={add}>Add</button>
            <button onClick={clear}>Clear</button>
            <button onClick={show}>Show</button>
          </div>
          <div className="main">
            <h3>Members</h3>
            <table border="1">
              <tbody>
                <tr>
                  <th>Name</th>
                  <th>Email</th>
                  <th>Registration date</th>
                </tr>
              </tbody>
              {members.map((row) => {
                 return <tbody><tr><td>{row.name}</td><td>{row.email}</td><td>{row.creation_time}</td></tr></tbody> 
              })}
            </table>
          </div>
        </div>
    </div>
  );
}

export default App;
