import React, { useEffect, useState } from 'react';
import InvoiceAnalytics from "./pages/InvoiceAnalytics"
import InvoiceQuarterly from './pages/InvoiceQuarterly';

const App = () => {
 return(
  <div>
  <InvoiceAnalytics />
  <InvoiceQuarterly />
</div>
 )
};

export default App;
