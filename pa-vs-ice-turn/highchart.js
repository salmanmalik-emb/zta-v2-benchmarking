// Data retrieved from https://www.vikjavev.no/ver/#2022-06-13,2022-06-14

Highcharts.chart('container', {
  chart: {
    type: 'spline',
    scrollablePlotArea: {
      minWidth: 600,
      scrollPositionX: 1
    }
  },
  title: {
    text: 'CPU vs Throughput',
  },

  xAxis: {
    title: {
      text: 'Throughput'
    },
   
  },
  yAxis: {
    title: {
      text: 'CPU'
    },
    tickInterval: 1
  },
  series: [{
    name: 'Relay Node (Coturn)',
    data: [[3.94, 0.18], [6.37, 0.2], [8.22, 0.25],  [9.46, 0.28], [12.4, 0.37], [35, 0.94]]

  }, {
    name: 'ZT-Connector',
    data: [[3.81, 0.4], [6.13, 0.5], [7.98, 0.66], [9.09, 0.78], [12.3, 1], [34.13, 2.57]]
  }, {
    name: 'PA-5250',
    data: [[3.8,7],[6,10],[8,12],[9,16],[12,17],[34,28]]
  }],
});