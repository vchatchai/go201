import React, {Component} from 'react'
import axios  from 'axios'

import EmployeeList from './EmployeeList'
import AddEmployee from './AddEmployee';


export default class EmployeeApp extends Component {
    constructor(props){
        super(props)
        this.state = {employees:[]};
        this.addEmployee = this.addEmployee.bind(this)
        this.Axios = axios.create({
            headers:{'content-type': 'application/json'}
        });
    }

    componentDidMount() { 
        let _this = this;
        this.Axios.get('/employees')
        .then(function(response) {
            _this.setState({employees: response.data})
        })
        .catch(function(error){})
    }

    addEmployee(employeeName) {
        let _this = this;
        console.log(employeeName)
        
        _this.Axios.post(
            '/employee/add', 
            {
                FirstName: employeeName
            }
        ).then(
            function(response){
                _this.setState({employees: response.data})
            }
        ).catch(function(error) {})
    }

    render() { 
        return (
            <div>
                <AddEmployee addEmployee={this.addEmployee} />
                <EmployeeList employees={this.state.employees} />
            </div>
        )
    }
}