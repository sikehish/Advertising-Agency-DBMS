import React from 'react';
import { ResponsiveContainer, PieChart, Pie, Tooltip, Legend } from 'recharts';

const DurationCategoryChart = ({ data }) => {
  // Define colors for each parameter
  const colors = ['#8884d8', '#82ca9d', '#ffc658'];

  const formatTooltip = (value, name, props) => {
    console.log(name , value, props)
    if (props.dataKey === 'numAds') {
      return `${value} ads`;
    } else if (props.dataKey === 'totalCost') {
      return `${value} Rs`;
    }
    return value;
  };

  return (
    <ResponsiveContainer width="80%" height={400}>
      <PieChart>
        <Pie
          data={data}
          dataKey="numAds"
          nameKey="durationCategory"
          cx="50%" cy="50%"
          outerRadius={120}
          fill={colors[0]} // Color for NumAds
        //   label
        />
        <Pie
          data={data}
          dataKey="totalCost"
          nameKey="durationCategory"
          cx="50%" cy="50%"
          outerRadius={150}
          innerRadius={130}
          fill={colors[1]} // Color for TotalCost
        //   label
        />

        <Tooltip formatter={formatTooltip} />
        {/* <Legend /> */}
      </PieChart>
    </ResponsiveContainer>
  );
};

export default DurationCategoryChart;
