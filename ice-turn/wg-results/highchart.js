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
      max: 100
    },
    series: [{
      name: 'Relay Node',
      data: [[0, 0], [0, 0], [0, 0], [31.7, 4.5], [35.73, 8.48], [49.73, 9.1], [175, 40.7], [257.3, 41.9], [288, 60.5], [350.66, 56], [586.66, 53.6], [613.33, 84], [636, 83]]
  
    }, {
      name: 'Connector',
      data: [[30.8, 10], [33.73, 26.5], [34.13, 21.3], [46.9, 29], [47.1, 21], [173.33, 58.7], [240, 81.4], [244, 74.6], [272, 78], [337.33, 92.1], [549.33, 98], [594.66, 98], [609.33, 93]]
    }],
  });