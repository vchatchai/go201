import React, {Component} from 'react'
import  Employee from './Employee'

export default class EmployeeList extends Component { 
    render() { 

        // console.log(this.props.employees)

        var employees = this.props.employees.map((employee,i)=> (<Employee key={i} employee={employee} />  ));

        return (
            <table>
                <tbody>
                    <tr>
                        <th>FirstName</th>
                        <th>LastName</th>
                    </tr>
                    {employees}
                </tbody>
            </table>
        )
    }
}