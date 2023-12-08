import React, { useEffect, useState } from 'react';
import { DataTable } from '../tables/DataTable';
import NumEmployeesRadarChart from '../charts/NumEmployeesRadarChart';

const NumEmployees = () => {
  const [numEmployeesData, setNumEmployeesData] = useState([]);

  useEffect(() => {
    fetch('api/v1/analytics/num-employees')
      .then(response => response.json())
      .then(data => setNumEmployeesData(data))
      .catch(error => console.error('Error fetching data:', error));
  }, []);

  return (
    <div className='flex m-10 justify-center items-center'>
      {numEmployeesData.length > 0 && <NumEmployeesRadarChart data={numEmployeesData} />}
      {numEmployeesData.length > 0 && <DataTable data={numEmployeesData} tableName={"Department Employees"}/>}
    </div>
  );
};

export default NumEmployees;