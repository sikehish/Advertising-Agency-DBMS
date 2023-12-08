import React, { useEffect, useState } from 'react';
import { DataTable } from '../tables/DataTable';
import SalaryExpenseRadialBarChart from '../charts/SalaryExpenseRadialBarChart';

const TotalSalary = () => {
  const [salaryData, setSalaryData] = useState([]);

  useEffect(() => {
    fetch('api/v1/analytics/total-salary')
      .then(response => response.json())
      .then(data => setSalaryData(data))
      .catch(error => console.error('Error fetching data:', error));
  }, []);

  return (<>
      {/* <h1>Salary Data Visualization</h1> */}
    <div className='flex m-10 justify-center items-center'>
      {salaryData.length > 0 && <SalaryExpenseRadialBarChart data={salaryData} />}
      {salaryData.length > 0 && <DataTable data={salaryData} tableName={"Department Salary Expenditure"}/>}
    </div>
  </>
  );
};

export default TotalSalary;