<template>
  <div :id="id" :class="className" :style="{height:height,width:width}"/>
</template>

<script>
import echarts from 'echarts'
import resize from './resize'
import {hashCode} from "@/utils";

export default {
  mixins: [resize],
  props: {
    className: {
      type: String,
      default: 'chart'
    },
    id: {
      type: String,
      default: 'chart'
    },
    width: {
      type: String,
      default: '200px'
    },
    height: {
      type: String,
      default: '200px'
    }
  },
  data() {
    return {
      chart: null,
      option: undefined,
    }
  },
  mounted() {
    this.initChart()
  },
  beforeDestroy() {
    if (!this.chart) {
      return
    }
    this.chart.dispose()
    this.chart = null
  },
  methods: {
    initChart() {
      this.chart = echarts.init(document.getElementById(this.id))
      const xData = (function () {
        const data = []
        for (let i = 0; i < 9; i++) {
          data.push('[ ' + i * 10 + ' , ' + (i + 1) * 10 + ' )')
        }
        data.push('[ 90 , 100 ]')
        return data
      }())
      this.option = {
        backgroundColor: '#344b58',
        title: {
          text: '成绩分析',
          x: '20',
          top: '20',
          textStyle: {
            color: '#fff',
            fontSize: '22'
          },
          subtextStyle: {
            color: '#90979c',
            fontSize: '16'
          }
        },
        tooltip: {},
        grid: {
          left: '5%',
          right: '5%',
          borderWidth: 0,
          top: 150,
          bottom: 95,
          textStyle: {
            color: '#fff'
          }
        },
        legend: {
          x: '5%',
          top: '10%',
          textStyle: {
            color: '#90979c'
          }
        },
        calculable: true,
        xAxis: [{
          type: 'category',
          name: '分数区间（左闭右开）',
          nameLocation: 'middle',
          nameGap: 50,
          nameTextStyle: {
            color: '#fff',
            fontSize: '18'
          },
          axisLine: {
            lineStyle: {
              color: '#90979c'
            }
          },
          splitLine: {
            show: false
          },
          axisTick: {
            show: false
          },
          splitArea: {
            show: false
          },
          axisLabel: {
            interval: 0,
            formatter: '{value} 分',
            textStyle: {
              color: '#f1f7fc'
            }
          },
          data: xData
        }],
        yAxis: [{
          type: 'value',
          name: '区间人数',
          nameLocation: 'middle',
          nameGap: 50,
          nameTextStyle: {
            color: '#fff',
            fontSize: '18'
          },
          minInterval: 1, // only integers for yAxis
          splitLine: {
            show: false
          },
          axisLine: {
            lineStyle: {
              color: '#90979c'
            },
          },
          axisTick: {
            show: false
          },
          axisLabel: {
            interval: 0,
            formatter: '{value} 人',
            textStyle: {
              color: '#f1f7fc'
            }
          },
          splitArea: {
            show: false
          }
        }]
      }
      this.chart.setOption(this.option)
    },
    setData(list) {
      this.chart.clear()
      let colors = ['Aquamarine', 'DeepSkyBlue', 'HotPink', 'NavajoWhite', 'RoyalBlue',
        'Tan', 'YellowGreen', 'SlateGray', 'SteelBlue', 'Orange', 'LightCyan']
      this.option.series = list.map(e => {
        return {
          name: e.name,
          type: 'bar',
          itemStyle: {
            normal: {
              color: colors[hashCode(e.name) % colors.length],
              barBorderRadius: 0,
              label: {
                show: true,
                position: 'top',
                formatter(p) {
                  return p.value > 0 ? p.value : ''
                }
              }
            }
          },
          data: e.data
        }
      })
      this.chart.setOption(this.option)
    }
  }
}
</script>
