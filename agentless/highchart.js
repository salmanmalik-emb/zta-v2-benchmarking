
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
  yAxis: [{
    labels: {
          format: '{value}%',
          style: {
              color: Highcharts.getOptions().colors[0]
          }
      },
    title: {
      text: 'CPU'
    },
    max: 100
  }],
  series: [{
    name: 'Relay Node',
    data: [[0, 2.6], [364.8, 23], [871.8, 56.6], [966.1, 63.6], [1372.8, 89.6], [1394.8, 86.6]]

  }, {
    name: 'Connector',
    data: [[0, 0.45], [359.3, 3.8], [869.1, 10.3], [963.4, 11.8], [1369.73, 16.81], [1388.73, 18.1]]
  }],
});