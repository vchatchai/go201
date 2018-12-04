import React, {Component} from 'react'

export default class Employee extends Component { 
    render() { 
        console.log(this.props.employee.FirstName)
        return (
            
            <tr>
                <td>{this.props.employee.FirstName}</td>
                <td>{this.props.employee.LastName}</td>
            </tr>
        )
    }
}