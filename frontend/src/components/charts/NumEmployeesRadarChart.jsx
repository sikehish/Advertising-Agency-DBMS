import React from 'react';
import { Radar, RadarChart, PolarGrid, PolarAngleAxis, PolarRadiusAxis, ResponsiveContainer } from 'recharts';

const NumEmployeesRadarChart = ({ data }) => {
  const departments = data.map((entry) => entry.department);
  const maxNumEmployees = Math.max(...data.map((entry) => entry.numEmployees));

  return (
    <ResponsiveContainer width="100%" height={400}>
      <RadarChart outerRadius={150} width="100%" height={400} data={data}>
        <PolarGrid />
        <PolarAngleAxis dataKey="department" />
        <PolarRadiusAxis angle={30} domain={[0, maxNumEmployees]} />
        <Radar name="Number of Employees" dataKey="numEmployees" stroke="#8884d8" fill="#8884d8" fillOpacity={0.6} />
      </RadarChart>
    </ResponsiveContainer>
  );
};

export default NumEmployeesRadarChart;
