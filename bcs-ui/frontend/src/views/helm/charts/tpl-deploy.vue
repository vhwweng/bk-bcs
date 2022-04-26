<template>
    <div class="biz-content">
        <div class="biz-top-bar">
            <div class="biz-helm-title">
                <a class="bcs-icon bcs-icon-arrows-left back" href="javascript:void(0);" @click="goToHelmIndex"></a>
                <span>{{ chartName }}</span>
            </div>
            <bcs-steps ext-cls="update-app-steps"
                :controllable="controllableSteps.controllable"
                :steps="controllableSteps.steps"
                :cur-step.sync="controllableSteps.curStep"
                @step-changed="stepChanged">
            </bcs-steps>
        </div>
        <div class="biz-content-wrapper" v-bkloading="{ isLoading: updateInstanceLoading, zIndex: 100 }">
            <div class="step-content">
                <!-- 基本信息 -->
                <div class="basic-content" v-show="controllableSteps.curStep === 1">
                    <div class="content-item">
                        <svg style="display: none;">
                            <title>{{$t('模板集默认图标')}}</title>
                            <symbol id="biz-set-icon" viewBox="0 0 32 32">
                                <path d="M6 3v3h-3v23h23v-3h3v-23h-23zM24 24v3h-19v-19h19v16zM27 24h-1v-18h-18v-1h19v19z"></path>
                                <path d="M13.688 18.313h-6v6h6v-6z"></path>
                                <path d="M21.313 10.688h-6v13.625h6v-13.625z"></path>
                                <path d="M13.688 10.688h-6v6h6v-6z"></path>
                            </symbol>
                        </svg>
                        <div class="logo-wapper" v-if="123 && isImage('')" @click="gotoHelmTplDetail">
                            <img :src="123" style="width: 100px;">
                        </div>
                        <svg class="logo-wapper" v-else>
                            <use xlink:href="#biz-set-icon"></use>
                        </svg>
                        
                        <span class="basic-wapper">
                            <div class="appName">{{ 123 }}</div>
                            <div class="desc">{{ $t('简介: ') }} <span class="">{{ 123 || '--' }}</span></div>
                        </span>
                    </div>
                    <div class="basic-box">
                        <div class="inner">
                            <div class="inner-item">
                                <label class="title">
                                    {{$t('名称')}}
                                    <bcs-popover :content="$t('Release名称只能由小写字母数字或者-组成')" placement="top">
                                        <span class="bk-badge">
                                            <i class="bcs-icon bcs-icon-question-circle f14"></i>
                                        </span>
                                    </bcs-popover>
                                </label>
                                <bkbcs-input :placeholder="$t('请输入Release名称')" />
                            </div>

                            <div class="inner-item">
                                <label class="title">{{$t('版本')}}</label>

                                <div style="display: flex;">
                                    <bk-selector
                                        :placeholder="$t('请选择')"
                                        style="display: inline-block; vertical-align: middle;"
                                        searchable
                                        :selected.sync="tplsetVerIndex"
                                        :list="curTplVersions"
                                        :disabled="isTplSynLoading"
                                        setting-key="version"
                                        display-key="version"
                                        search-key="version"
                                        @item-selected="getTplDetail">
                                    </bk-selector>

                                    <bk-button class="bk-button bk-default is-outline is-icon" v-bk-tooltips.top="$t('同步仓库')" @click="syncHelmTpl">
                                        <div class="bk-spin-loading bk-spin-loading-mini bk-spin-loading-default" style="margin-top: -3px;" v-if="isTplSynLoading">
                                            <div class="rotate rotate1"></div>
                                            <div class="rotate rotate2"></div>
                                            <div class="rotate rotate3"></div>
                                            <div class="rotate rotate4"></div>
                                            <div class="rotate rotate5"></div>
                                            <div class="rotate rotate6"></div>
                                            <div class="rotate rotate7"></div>
                                            <div class="rotate rotate8"></div>
                                        </div>
                                        <i class="bcs-icon bcs-icon-refresh" v-else></i>
                                    </bk-button>
                                </div>
                            </div>
                        </div>
                        <div class="inner">
                            <div class="inner-item">
                                <label class="title">{{$t('所属集群')}}</label>
                                <bk-selector
                                    :placeholder="$t('请选择')"
                                    :searchable="true"
                                    :selected.sync="curClusterId"
                                    :field-type="'cluster'"
                                    :list="clusterList"
                                    :setting-key="'cluster_id'"
                                    :display-key="'name'"
                                    :disabled="!!globalClusterId">
                                </bk-selector>
                            </div>
                            <div class="inner-item">
                                <label class="title">{{$t('命名空间')}}</label>
                                <div style="display: flex;align-items: center;">
                                    <bcs-select style="width: 100%;"
                                        searchable
                                        :clearable="false"
                                        v-model="namespaceId">
                                        <bcs-option v-for="(item, index) in namespaceList"
                                            :key="item.id"
                                            :id="item.id"
                                            :name="item.name"
                                            v-authority="{
                                                clickable: webAnnotations.perms[item.iam_ns_id]
                                                    && webAnnotations.perms[item.iam_ns_id].namespace_scoped_use,
                                                actionId: 'namespace_scoped_use',
                                                resourceName: item.name,
                                                disablePerms: true,
                                                permCtx: {
                                                    project_id: projectId,
                                                    cluster_id: item.cluster_id,
                                                    name: item.name
                                                }
                                            }"
                                            @click.native="getClusterInfo(index, item)">
                                        </bcs-option>
                                        <div slot="extension" style="cursor: pointer;"
                                            @click="goNamespaceList">
                                            <i class="bcs-icon bcs-icon-apps"></i>
                                            <span style="font-size: 14px">{{$t('命名空间列表')}}</span>
                                        </div>
                                    </bcs-select>
                                    <i v-bk-tooltips.top="$t('如果Chart中已经配置命名空间，则会使用Chart中的命名空间，会导致不匹配等问题;建议Chart中不要配置命名空间')" class="bcs-icon bcs-icon-question-circle f14 ml5"></i>
                                </div>
                            </div>
                            <p class="biz-tip pt10" id="cluster-info" style="clear: both;" v-if="clusterInfo" v-html="clusterInfo"></p>
                        </div>

                        <div class="desc-inner">
                            <label class="title">{{$t('描述')}}</label>
                            <bcs-input type="textarea" :rows="3"></bcs-input>
                        </div>

                        <div class="readme-inner">
                            <label class="title">{{$t('自述')}}</label>
                            <div class="readme-content">
                                Rancher Alerting Drivers This chart enables ability to capture backups of the Rancher application and restore from these backups. This chart can be used to migrate Rancher from one Kubernetes cluster to a different Kubernetes cluster.For more information on how to use the feature,
                                refer to our docs. This chart installs one or more Alertmanager Webhook Receiver Integrations (i.e. Drivers).
                            </div>
                        </div>
                    </div>
                </div>

                <!-- 配置信息 -->
                <div class="config-content" v-show="controllableSteps.curStep === 2">
                    <bk-tab type="border-card" active-name="chart" class="mt20">
                        <bk-tab-panel name="chart" :title="$t('Chart配置选项')">
                            <div slot="content" style="min-height: 180px;">
                                <section class="value-file-wrapper">
                                    {{$t('Values文件：')}}
                                    <bk-selector
                                        style="width: 300px;"
                                        :placeholder="$t('请选择')"
                                        :searchable="true"
                                        :selected.sync="curValueFile"
                                        :list="curValueFileList"
                                        :setting-key="'name'"
                                        :display-key="'name'"
                                        @item-selected="changeValueFile">
                                    </bk-selector>
                                    <bcs-popover placement="top">
                                        <span class="bk-badge ml5">
                                            <i class="bcs-icon bcs-icon-question-circle f14"></i>
                                        </span>
                                        <div slot="content">
                                            <p>{{ $t('Values文件包含两类:') }}</p>
                                            <p>{{ $t('- 以values.yaml结尾，例如xxx-values.yaml文件') }}</p>
                                            <p>{{ $t('- bcs-values目录下的文件') }}</p>
                                        </div>
                                    </bcs-popover>
                                </section>
                                <bk-tab
                                    :type="'fill'"
                                    :size="'small'"
                                    :active-name.sync="curEditMode"
                                    class="biz-tab-container"
                                    @tab-changed="helmModeChangeHandler">
                                    <bk-tab-panel name="yaml-mode" :title="$t('YAML模式')">
                                        <div style="width: 100%; min-height: 600px;">
                                            <p class="biz-tip p15" style="color: #63656E; overflow: hidden;">
                                                <i class="bcs-icon bcs-icon-info-circle biz-warning-text mr5"></i>
                                                {{$t('YAML初始值为创建时Chart中values.yaml内容，后续更新部署以该YAML内容为准，YAML内容最终通过`--values`选项传递给`helm template`命令')}}
                                            </p>
                                            <div v-bkloading="{ isLoading: isSyncYamlLoading, color: '#272822' }">
                                                <ace
                                                    v-if="curEditMode === 'yaml-mode'"
                                                    :value="curTplYaml"
                                                    :width="yamlConfig.width"
                                                    :height="yamlConfig.height"
                                                    :lang="yamlConfig.lang"
                                                    :read-only="yamlConfig.readOnly"
                                                    :full-screen="yamlConfig.fullScreen"
                                                    @init="editorInit">
                                                </ace>
                                            </div>
                                        </div>
                                    </bk-tab-panel>
                                    <bk-tab-panel name="form-mode" :title="$t('表单模式')">
                                        <p class="biz-tip p15" style="color: #63656E;">
                                            <i class="bcs-icon bcs-icon-info-circle biz-warning-text mr5"></i>{{$t('表单根据Chart中questions.yaml生成，表单修改后的数据会自动同步给YAML模式')}}
                                        </p>
                                        <template v-if="true">
                                            <bk-form-creater :form-data="formData" ref="bkFormCreater"></bk-form-creater>
                                        </template>
                                        <template v-else>
                                            <div class="biz-guard-box" v-if="!isQuestionsLoading">
                                                <span>{{$t('您可以参考')}}
                                                    <a class="bk-text-button" :href="PROJECT_CONFIG.doc.questionsYaml" target="_blank">{{$t('指引')}}</a>
                                                    {{$t('通过表单模式配置您的Helm Release 参数')}}，
                                                </span>
                                                <span>{{$t('也可以通过')}}<a href="javascript:void(0)" class="bk-text-button" @click="editYaml"></a>{{$t('直接修改Helm Release参数')}}</span>
                                            </div>
                                        </template>
                                    </bk-tab-panel>
                                </bk-tab>
                            </div>
                        </bk-tab-panel>
                        <bk-tab-panel name="helm" :title="$t('Helm部署选项')">
                            <div class="helm-set-panel">
                                <ul class="mt10">
                                    <!-- 常用枚举项 -->
                                    <li v-for="command of commandList" :key="command.id">
                                        <bk-checkbox
                                            class="mr5"
                                            v-model="helmCommandParams[command.id]" />
                                        <span class="mb5" style="display: inline-block;">
                                            {{command.desc}}
                                            <i style="font-size: 12px;cursor: pointer;"
                                                class="bcs-icon bcs-icon-info-circle"
                                                v-bk-tooltips.top="command.id" />
                                        </span>
                                    </li>
                                    <li class="mt10">
                                        <div style="margin-bottom:4px;">
                                            {{ $t('超时时间') }}
                                            <i style="font-size: 12px;cursor: pointer;"
                                                class="bcs-icon bcs-icon-info-circle"
                                                v-bk-tooltips.top="'--timeout'" />
                                        </div>
                                        <bk-input
                                            v-model="timeoutValue"
                                            placeholder="500"
                                            style="width: 200px; margin-right:4px;" />
                                        <span>{{ $t('秒') }}</span>
                                    </li>
                                    <!-- 高级选项 -->
                                    <button class="bk-text-button f12 mb10 pl0 mt10" @click.stop.prevent="toggleHign">
                                        {{$t('高级设置')}}<i class="bcs-icon bcs-icon-angle-double-down ml5"></i>
                                        <i style="font-size: 12px; cursor: pointer;"
                                            class="bcs-icon bcs-icon-info-circle ml5"
                                            v-bk-tooltips.top="hignDesc" />
                                    </button>
                                    <div v-show="isHignPanelShow">
                                        <div class="biz-key-value-wrapper mb10">
                                            <li class="biz-key-value-item mb10" v-for="(item, index) in hignSetupMap" :key="index">
                                                <bk-input style="width: 280px;" v-model="item.key" @change="handleHignkeyChange(item.key ,index)" />
                                                <span class="equals-sign">=</span>
                                                <bk-input style="width: 280px;" :placeholder="$t('值')" v-model="item.value" />
                                                <button class="action-btn" @click.stop.prevent>
                                                    <i class="bk-icon icon-plus-circle mr5" v-if="index === 0" @click.stop.prevent="addHign"></i>
                                                    <i class="bk-icon icon-minus-circle" @click.stop.prevent="delHign(index)"></i>
                                                </button>
                                                <p class="error-key" v-if="item.errorKeyTip">{{ item.errorKeyTip }}</p>
                                            </li>
                                        </div>
                                    </div>
                                </ul>
                            </div>
                        </bk-tab-panel>
                    </bk-tab>
                </div>
            </div>
        </div>
        <div class="biz-footer-actions">
            <bcs-button theme="primary" v-if="controllableSteps.curStep === 2" @click="controllableSteps.curStep = 1">{{ $t('上一步') }}</bcs-button>
            <bcs-button theme="primary" v-else @click="controllableSteps.curStep = 2">{{ $t('下一步') }}</bcs-button>
            <bcs-button type="primary" v-if="controllableSteps.curStep = 2" @click="createApp">{{$t('部署')}}</bcs-button>
            <bcs-button type="default" @click="showPreview"> {{$t('预览')}} </bcs-button>
            <bcs-button type="default" @click="goBack">{{$t('取消')}}</bcs-button>
        </div>
        
        <!-- 预览 -->
        <bk-sideslider
            :is-show.sync="previewEditorConfig.isShow"
            :title="previewEditorConfig.title"
            :quick-close="true"
            :width="1000">
            <div slot="content" :style="{ height: `${winHeight - 70}px` }" v-bkloading="{ isLoading: previewInstanceLoading }">
                <!-- <template v-if="tplPreviewList.length">
                    <div class="biz-resource-wrapper" style="height: 100%; flex: 1;">
                        <resizer :class="['resize-layout fl']"
                            direction="right"
                            :handler-offset="3"
                            :min="250"
                            :max="400">
                            <div class="tree-box">
                                <bcs-tree
                                    ref="tree1"
                                    :data="treeData"
                                    :node-key="'id'"
                                    :has-border="true"
                                    @on-click="getFileDetail">
                                </bcs-tree>
                            </div>
                        </resizer>

                        <div class="resource-box">
                            <div class="biz-code-wrapper" style="height: 100%;">
                                <ace
                                    :value="curReourceFile.value"
                                    :width="editorConfig.width"
                                    :height="editorConfig.height"
                                    :lang="editorConfig.lang"
                                    :read-only="editorConfig.readOnly"
                                    :full-screen="editorConfig.fullScreen">
                                </ace>
                            </div>
                        </div>
                    </div>
                </template>
                <bcs-exception type="empty" scene="part" v-else></bcs-exception> -->
            </div>
        </bk-sideslider>

        <!-- 错误弹框 -->
        <bk-dialog
            :is-show.sync="errorDialogConf.isShow"
            :width="750"
            :has-footet="false"
            :title="errorDialogConf.title"
            @cancel="hideErrorDialog">
            <template slot="content">
                <div class="bk-intro bk-danger pb30 mb15" v-if="errorDialogConf.message" style="position: relative;">
                    <pre class="biz-error-message">
                        {{errorDialogConf.message}}
                    </pre>
                    <bk-button size="small" type="default" id="error-copy-btn" :data-clipboard-text="errorDialogConf.message"><i class="bcs-icon bcs-icon-clipboard mr5"></i>{{$t('复制')}}</bk-button>
                </div>
                <div class="biz-message" v-if="errorDialogConf.errorCode === 40031">
                    <h3>{{$t('您需要')}}：</h3>
                    <p>1、{{$t('在集群页面，启用Helm')}}</p>
                    <p>2、{{$t('或者联系')}}【<a :href="PROJECT_CONFIG.doc.contact" class="bk-text-button">{{$t('蓝鲸容器助手')}}</a>】</p>
                </div>
                <div class="biz-message" v-else-if="errorDialogConf.actionType === 'previewApp'">
                    <h3>{{$t('您可以')}}：</h3>
                    <p>1、{{$t('检查Helm Chart是否存在语法错误')}}</p>
                    <p>2、{{$t('前往Helm Release列表页面，更新Helm Release')}}</p>
                </div>
                <div class="biz-message" v-else>
                    <h3>{{$t('您可以')}}：</h3>
                    <p>1、{{$t('更新Helm Chart，并推送到项目Chart仓库')}}</p>
                    <p>2、{{$t('前往Helm Release列表页面，更新Helm Release')}}</p>
                </div>
            </template>
            <div slot="footer">
                <div class="biz-footer">
                    <bk-button type="primary" @click="hideErrorDialog">{{$t('知道了')}}</bk-button>
                </div>
            </div>
        </bk-dialog>
    </div>
</template>

<script>
    import { onMounted, ref, reactive, computed, watch } from "@vue/composition-api"
    import baseMixin from '@/mixins/helm/mixin-base'
    // import resizer from '@/components/resize'
    // import MarkdownIt from 'markdown-it'
    // import Clipboard from 'clipboard'
    // import yamljs from 'js-yaml'
    // import path2tree from '@/common/path2tree'

    export default {
        name: 'tplDeploy',
        components: {
            // resizer
        },
        mixins: [baseMixin],
        setup (props, ctx) {
            const { $i18n, $route, $store, $INTERNAL } = ctx.root
            const appName = ref('')
            const curClusterId = ref('')
            const tplsetVerIndex = ref('')
            const isTplSynLoading = ref(false)
            const namespaceList = ref([])
            const winHeight = ref(0)
            const curTplVersions = ref([])
            const curTpl = ref({
                data: {
                    name: ''
                }
            })
            const isTplVerLoading = ref(false)
            const webAnnotations = reactive({
                perms: {}
            })
            const previewEditorConfig = reactive({
                isShow: false,
                title: $i18n.t('预览'),
                width: '100%',
                height: '100%',
                lang: 'yaml',
                readOnly: true,
                fullScreen: false,
                value: '',
                editors: []
            })

            const errorDialogConf = reactive({
                title: '',
                isShow: false,
                message: '',
                errorCode: 0
            })
            const controllableSteps = reactive({
                controllable: true,
                steps: [
                    { title: '基本信息', icon: 1 },
                    { title: '配置信息', icon: 2 }
                ],
                curStep: 1
            })

            const projectId = computed(() => {
                return $route.params.projectId
            })
            // const projectCode = computed(() => {
            //     return $route.params.projectCode
            // })
            const chartName = computed(() => {
                return $route.params.chartName
            })
            const tplList = computed(() => {
                return $store.state.helm.tplList
            })
            const globalClusterId = computed(() => {
                return $store.state.helm.curClusterId
            })
            const tplId = computed(() => {
                return $route.params.tplId
            })
            const clusterList = computed(() => {
                return $store.state.helm.clusterList
            })

            watch(globalClusterId, (val) => {
                curClusterId.value = val
            }, {
                immediate: true
            })

            watch(curClusterId, () => {
                getNamespaceList(tplId.value)
            })

            /**
             * 获取命名集群和空间列表
             */
            const getNamespaceList = async (chartId) => {
                if (!curClusterId.value) return

                const res = await $store.dispatch(
                    'helm/getNamespaceList',
                    {
                        projectId: projectId.value,
                        params: {
                            chart_id: chartId,
                            cluster_id: curClusterId.value
                        }
                    }
                ).catch(() => false)

                if (!res) return

                namespaceList.value = res.data
                webAnnotations.value = res.web_annotations || { perms: {} }
            }

            /**
             * 点击修改步骤
             * @param {Boolean} val 当前步骤
             */
            const stepChanged = (val) => {
                controllableSteps.curStep = val
            }

            /**
             * 获取模板版本列表
             */
            const getTplVersions = async () => {
                if ($INTERNAL) {
                    // 内部版本
                    const tplId = curTpl.value.name
                    const isPublic = curTpl.value.repository.name === 'public-repo'
                    const res = await $store.dispatch('helm/getTplVersionList', {
                        projectId: projectId.value,
                        tplId,
                        isPublic
                    })
                    curTplVersions.value = res.data
                } else {
                    // 外部版本
                    const tplId = curTpl.value.id
                    const res = await $store.dispatch('helm/getTplVersions', {
                        projectId: projectId.value,
                        tplId
                    })
                    curTplVersions.value = res.data.results || []
                }
                console.log(curTplVersions, '11111111')
                
                setTimeout(() => {
                    isTplVerLoading.value = false
                }, 1000)
            }

            /**
             * 获取模板
             * @param  {number} id 模板ID
             * @return {object} result 模板
             */
            const getTplById = async (id) => {
                let list = tplList.value

                // 如果没有缓存，获取远程数据
                if (!list.length) {
                    const projectId = this.projectId
                    const res = await this.$store.dispatch('helm/asyncGetTplList', projectId)
                    if (!res) return
                    list = res.data
                }

                const result = list.find(item => item.id === Number(id))
                return result || {}
            }

            const isQuestionsLoading = ref(false)
            const formData = ref({})
            const curVersionData = ref({})
            const curValueFileList = ref([])
            const curTplReadme = ref('')
            const curTplYaml = ref('')
            const yamlFile = ref('')
            /**
             * 根据版本号获取模板详情
             * @param  {number} index 索引
             * @param  {object} data 数据
             */
            const getTplDetail = async (index, data) => {
                const list = []
                const version = index
                const versionId = curTplVersions.value.find(item => item.version === index).id
                const isPublic = curTpl.value.repository.name === 'public-repo'

                isQuestionsLoading.value = true
                try {
                    const fnPath = $INTERNAL ? 'helm/getChartVersionDetail' : 'helm/getChartByVersion'
                    const res = await $store.dispatch(fnPath, {
                        projectId: projectId.value,
                        chartId: $INTERNAL ? curTpl.value.name : curTpl.value.id,
                        version: $INTERNAL ? version : versionId,
                        isPublic
                    })
                    const tplData = res.data
                    const files = res.data.data.files
                    const tplName = tplData.name
                    const bcsTplName = tplData.name + '/bcs-values'
                    const regex = new RegExp(`^${tplName}\\/[\\w-]*values.(yaml|yml)$`)
                    const bcsRegex = new RegExp(`^${bcsTplName}\\/[\\w-]*.(yaml|yml)$`)

                    formData.value = res.data.data.questions
                    curVersionData.value = tplData

                    for (const key in files) {
                        if (bcsRegex.test(key)) {
                            const catalog = key.split('/')
                            const fileName = catalog[catalog.length - 2] + '/' + catalog[catalog.length - 1]
                            list.push({
                                name: fileName,
                                content: files[key]
                            })
                        }
                        if (regex.test(key)) {
                            const catalog = key.split('/')
                            const fileName = catalog[catalog.length - 1]
                            list.push({
                                name: fileName,
                                content: files[key]
                            })
                        }
                    }
                    curValueFileList.value.splice(0, curValueFileList.value.length, ...list)
                    curTplReadme.value = files[`${tplName}/README.md`]
                    curTplYaml.value = files[`${tplName}/values.yaml`]
                    yamlFile.value = files[`${tplName}/values.yaml`]
                    editYaml()
                    curTpl.value.description = res.data.data.description
                } catch (e) {
                } finally {
                    isQuestionsLoading.value = false
                }
            }

            const curEditMode = ref('')
            const isSyncYamlLoading = ref(false)
            const bkFormCreater = ref(null)
            const yamlConfig = reactive({
                isShow: false,
                title: $i18n.t('预览'),
                width: '100%',
                height: '700',
                lang: 'yaml',
                readOnly: false,
                fullScreen: false,
                value: '',
                editors: []
            })

            /**
             * 编辑yamml
             */
            const editYaml = async () => {
                curEditMode.value = 'yaml-mode'
                let formData = []

                isSyncYamlLoading.value = true
                // 将数据配置的数据和yaml的数据进行合并同步
                if (bkFormCreater.value) {
                    formData = bkFormCreater.value.getFormData()
                }

                yamlConfig.isShow = true

                const res = await $store.dispatch('helm/syncJsonToYaml', {
                    json: formData,
                    yaml: curTplYaml.value
                }).catch(() => false)
                isSyncYamlLoading.value = false
                if (!res) return

                curTplYaml.value = res.data.yaml
            }

            onMounted(async () => {
                curTpl.value = await getTplById(tplId.value)
                getTplVersions()
                getNamespaceList(tplId.value)
                winHeight.value = window.innerHeight
            })

            return {
                chartName,
                appName,
                tplsetVerIndex,
                curEditMode,
                isSyncYamlLoading,
                isTplSynLoading,
                bkFormCreater,
                yamlConfig,
                curTplVersions,
                previewEditorConfig,
                errorDialogConf,
                controllableSteps,
                tplList,
                globalClusterId,
                clusterList,
                stepChanged,
                getTplDetail
            }
        }
    }
</script>

<style lang="postcss" scoped>
    @import '../common.css';
    .biz-helm-title {
        display: inline-block;
        height: 60px;
        line-height: 60px;
        font-size: 16px;
        margin-left: 20px;
    }
    .update-app-steps {
        width: 300px;
        position: absolute;
        right: 26px;
        top: 18px;
    }
    .step-content {
        background-color: #fff;
        padding: 24px 32px;
    }
    .biz-footer-actions {
        position: fixed;
        bottom: 0;
        width: 100%;
        height: 60px;
        background-color: #fff;
        padding: 14px 24px;
        z-index: 20;
    }
    .basic-content {
        .content-item {
            width: 800px;
            display: flex;
            .logo-wapper {
                display: inline-block;
                width: 60px;
                height: 60px;
                margin-right: 24px;
                vertical-align: middle;
                fill: #ebf0f5;
            }
            .basic-wapper {
                display: inline-block;
                .appName {
                    margin-top: 6px;
                    font-size: 16px;
                    color: #313238;
                }
                .desc {
                    margin-top: 10px;
                    color: #313238;
                    font-size: 12px;
                    span {
                        color: #979ba5;
                    }
                }
            }
        }
        .basic-box {
            margin-top: 24px;
            .inner {
                display: -webkit-box;
                width: 783px;
            }
            .inner-item {
                width: 360px;
                margin-right: 64px;
                margin-bottom: 24px;
            }
            .title {
                display: inline-block;
                margin-bottom: 6px;
                font-size: 12px;
            }
            .desc-inner {
                width: 783px;
                margin-bottom: 24px;
            }
            .readme-inner {
                .readme-content {
                    width: 783px;
                    margin-bottom: 20px;
                    background: #f0f1f5;
                    border: 1px solid #dcdee5;
                    border-radius: 2px;
                    font-size: 12px;
                    padding: 14px 8px;
                }
            }
        }
    }

    .config-content {
        .value-file-wrapper {
            padding: 15px 16px;
            border: 1px solid #dcdee5;
            border-radius: 2px;
            background: #f9fbfd;
            position: relative;
            bottom: -1px;
            font-size: 13px;
        }
        .action-btn {
            width: auto;
            padding: 0;
            height: 30px;
            text-align: center;
            display: inline-block;
            border: none;
            background: transparent;
            outline: none;
            margin-left: 5px;

            .bk-icon {
                width: 24px;
                height: 24px;
                line-height: 24px;
                border-radius: 50%;
                vertical-align: middle;
                color: #999999;
                font-size: 24px;
                display: inline-block;

                &:hover {
                    color: $primaryColor;
                    border-color: $primaryColor;
                }
            }
        }
    }
    .difference-code {
        height: 350px;
    }

    .editor-header {
        display: flex;
        display: flex;
        background: #eee;
        border-radius: 2px 2px 0 0;

        > div {
            padding: 5px;
            width: 50%;
        }
    }
    .biz-error-message {
        white-space: pre-line;
        text-align: left;
        max-height: 200px;
        overflow: auto;
        margin: 0;
    }
    .biz-message {
        margin-bottom: 0;

        h3 {
            text-align: left;
            font-size: 14px;
        }

        p {
            text-align: left;
            font-size: 13px;
        }
    }
    .biz-footer {
        text-align: right;
        padding-right: 15px;
    }
    #error-copy-btn {
        position: absolute;
        bottom: 5px;
        right: 5px;
    }
    .helm-set-panel {
        padding: 20px;
        li {
            font-size: 14px;
        }
    }
</style>
