import React, { useEffect, useState } from 'react';
import AdvertisementDurationStatsChart from '../charts/AdvertisementDurationStatsChart';
import { AdsDurationTable } from '../tables/AdsDurationTable';
import { DataTable } from '../tables/DataTable';

const AdvertisementDuration = () => {
  const [durationStats, setDurationStats] = useState([]);

  useEffect(() => {
    fetch('/api/v1/analytics/ads-duration')
      .then(response => response.json())
      .then(data => setDurationStats(data));
  }, []);

  return (<>
      {/* <h1>Advertisement Duration Stats {durationStats.length>0}</h1> */}
      <div className='flex m-10 justify-center items-center'>
      {durationStats.length > 0 && <AdvertisementDurationStatsChart data={durationStats} />}
      {durationStats.length > 0 && <DataTable data={durationStats} tableName={"Advertisement Duration(days) related data"}/>}
    </div>
  </>
  );
};

export default AdvertisementDuration;
