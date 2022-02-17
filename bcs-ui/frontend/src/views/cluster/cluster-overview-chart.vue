<template>
    <div v-bkloading="{ isLoading: showLoading }">
        <canvas :id="`${chartType}Chart`"></canvas>
    </div>
</template>

<script>
    import moment from 'moment'
    import BKChart from '@tencent/bkchart.js'

    export default {
        props: {
            chartType: {
                type: String
            },
            showLoading: {
                type: Boolean,
                default: true
            },
            data: {
                type: Array,
                default: () => []
            },
            // matrix，数据为二维数组
            // vector，数据为一维数组
            resultType: {
                type: String,
                default: 'matrix'
            }
        },
        data () {
            return {
                chart: null
            }
        },
        computed: {
            curCluster () {
                return this.$store.state.cluster.curCluster
            },
            isEn () {
                return this.$store.state.isEn
            }
        },
        watch: {
            data (v) {
                setTimeout(() => {
                    this.initData()
                }, 0)
            }
        },
        mounted () {
            if (!this.showLoading) {
                this.initData()
            }
            window.addEventListener('resize', this.resizeHandler)
        },
        destroyed () {
            window.removeEventListener('resize', this.resizeHandler)
        },
        methods: {
            initData () {
                let x = []
                const y = []
                const data = this.data.length ? this.data : [{ values: [[parseInt(String(+new Date()).slice(0, 10), 10), '0']] }]
                data.forEach(item => {
                    item.values.forEach(value => {
                        x.push(value[0])
                        y.push(value[1])
                    })
                })
                x = x.map(i => {
                    if (String(parseInt(i, 10)).length === 10) {
                        i = parseInt(i, 10) + '000'
                    }
                    return moment(parseInt(i, 10)).format('HH:mm')
                })
                this.initChart(x, y)
            },

            initChart (x, y) {
                let borderColor = 'rgb(48, 216, 120)'
                let backgroundColor = 'rgba(48, 216, 120, 0.2)'
                if (this.chartType === 'mem') {
                    borderColor = 'rgb(58, 132, 255)'
                    backgroundColor = 'rgba(58, 132, 255, 0.2)'
                }
                if (this.chartType === 'disk') {
                    borderColor = 'rgb(133, 60, 255)'
                    backgroundColor = 'rgba(133, 60, 255, 0.2)'
                }
                let tipsLabel = ''
                if (this.chartType === 'cpu') {
                    tipsLabel = this.$t('CPU使用率')
                }

                if (this.chartType === 'mem') {
                    tipsLabel = this.$t('内存使用率')
                }

                if (this.chartType === 'disk') {
                    tipsLabel = this.$t('磁盘使用率')
                }
                const ymd = moment().format('YYYY-MM-DD')
                const s = moment().format('ss')
                const chartId = `${this.chartType}Chart`
                const context = document.getElementById(chartId)
                this.chart = new BKChart(context, {
                    type: 'line',
                    data: {
                        labels: x,
                        datasets: [
                            {
                                label: false,
                                fill: true,
                                borderColor,
                                backgroundColor,
                                lineTension: 0.3,
                                borderWidth: 2,
                                pointRadius: 0,
                                data: y,
                                tipsLabel
                            }
                        ]
                    },
                    options: {
                        tickWidth: 20,
                        layout: {
                            padding: {
                                left: 20,
                                right: 20
                            }
                        },
                        interaction: {
                            intersect: false,
                            mode: 'index'
                        },
                        plugins: {
                            legend: {
                                display: false
                            },
                            crosshair: {
                                enabled: true,
                                style: {
                                    x: {
                                        enabled: true,
                                        borderStyle: 'solid',
                                        weight: 1
                                    },
                                    y: {
                                        enabled: false
                                    }
                                }
                            },
                            tooltip: {
                                displayColors: false,
                                callbacks: {
                                    title (data) {
                                        const label = data[0].dataset.tipsLabel
                                        const hm = data[0].label
                                        return `${ymd} ${hm}:${s}\n${label}：${parseFloat(data[0].raw).toFixed(2)}%`
                                    },
                                    label (item) {
                                        return item.dataset.label
                                    }
                                }
                            }
                        },
                        scales: {
                            x: {
                                min: 0,
                                ticks: {
                                    maxTicksLimit: 6,
                                    maxRotation: 0
                                }
                            },
                            y: {
                                min: 0,
                                ticks: {
                                    callback: function (value) {
                                        if (value >= 0) return value + '.0%'
                                        return value + '%'
                                    }
                                }
                            }
                        }
                    }
                })
            },
            resizeHandler () {
                this.chart && this.chart.resize()
            }
        }
    }
</script>
<style lang="postcss">
    .echarts {
        width: 100%;
        height: 250px;
    }
</style>
