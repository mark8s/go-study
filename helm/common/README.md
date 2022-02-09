# Common

**Common** 是一个公共的库类型chart，里面定义了一些公共的函数。它不能单独install，只能通过引用去使用它里面预定义的函数。

## 使用common

修改需要引入common的helm charts 的Chart.yaml文件，添加如下代码:

```yaml
dependencies:
  - name: common
    version: "1.0.0"
    repository: "file://../../../common"
```

假如，现在引入common的charts的名称为test。现在需要更新test的依赖，执行如下命令：

`helm dependency update test`

更新后，在test的templates目录下就可以使用common的函数了。如，我们直接使用里面的deployment函数，生成一个完整的deployment资源。

先创建一个deployment.yaml文件，然后放置以下代码

```gotemplate
{{- include "common.deployment" . }}
```

配合test的values.yaml中的内容，就可以生成一个deployment资源了。

## 函数

### common中包含的函数

| Name                      | Description                                      |
| ------------------------- | -----------------------------------------------  |
| `common.deployment`    | 生成deployment资源                     |
| `common.service` | 生成service资源 |
| `common.rbac.serviceAccount` | 生成serviceAccount资源 |
| `common.labels.standard` | 标准label |
| `common.labels.matchLabels` | matchLabel | 
| `common.names.fullname` | 名称 | 
| ... | | 

## Values文件规范

values.yaml文件中的内容，本质上还是为k8s的一些常见的资源去服务的，如deployment、service、configmap，所以values里面定义的一些变量也与k8s相关。
下面，列举一些values.yaml常见的内容，可以供你后续自定义charts时作为参考。

以下内容是使用common中deployment函数时，values.yaml文件中字段的标准写法！

#### replicaCount
定义replicaCount去描述资源使用的副本数，如：
```yaml
replicaCount: 1
```
#### image
定义image去描述资源使用的镜像，如：
```yaml
image:
  name: nginx
  pullPolicy: IfNotPresent
```
image.name: 代表一个完整的镜像地址。如果值是引用其他的变量值，需要加双引号，如 name: `"{{- .Values.global.dockerRegisty }}/{{- .Values.global.orchorImage }}"`
image.pullPolicy: 定义镜像pull策略，可不设置，k8s默认策略就是IfNotPresent。

#### fullnameOverride

定义资源名称，common的模板或者helm chart默认生成的_helper.tpl模板中，资源名称默认是使用的releaseName。如果想更改，可以在values中设置fullnameOverride。
如果是使用common中的deployment模板，并且使用helm一次性部署多个chart时，需要明确定义这个字段，因为deployment模板在fullnameOverride为空时，默认取release值，如果是一次性部署多个chart会出现资源名称冲突问题。
如：
```yaml
fullnameOverride: "demo"
```

#### env
定义资源中使用的环境变量，env应该定义成一个数组，如：
```yaml
env:
  - name: DISABLE_HTTP2
    value: "1"
  - name: NODE_IP
    valueFrom:
      fieldRef:
        apiVersion: v1
        fieldPath: status.hostIP
  - name: APIServerEndpoint
    value: turtle:16020
```

#### service
定义k8s service资源，如：
```yaml
service:
  #type: ClusterIP
  ports:
    - name: tcp-port-1400
      port: 1400
      protocol: TCP
      targetPort: 1400
```

#### serviceAccount
定义serviceAccount是否创建，如：
```yaml
serviceAccount:
  create: true
```

#### resources
定义pod中resource，如：
```yaml
resources:
  limits: {}
  requests:
    memory: 512Mi
    cpu: 300m
```

#### restartPolicy
定义重启策略，如：
```yaml
restartPolicy: Always
```

#### command
定义command

#### args
定义args，如：
```yaml
args:
    - "-zone={{- .Values.global.dnsa.args.zones }}"
    - -mode=master
    - -dnsa=dnsa-master:1401
    - "-mongo={{- .Values.global.dnsa.args.mongostr }}"
    - -dynamicheck=true
    - "-enablehttps={{ .Values.global.enable_https }}"
```

#### updateStrategy
定义更新策略，如：
```yaml
updateStrategy:
  type: RollingUpdate
```

#### commonLabels
定义标准labels，也就是k8s资源中metadata下面的labels，如：
```yaml
commonLabels:
  app: demo
```

#### commonAnnotations

定义标准annotations，也就是k8s资源中metadata下面的annotations，如：
```yaml
commonAnnotations:
  creater: mark
```

#### matchLables
定义deployment(或sts、ds)匹配pod的matchLables和pod的labels，如：

```yaml
matchLabels:
  app: test
```

#### podAnnotations
定义配pod的annotations，如：
```yaml
podAnnotations:
  creater: mark
```

#### containerPort
定义deployment等资源中的容器端口，如：
```yaml
containerPort:
  - name: http
    containerPort: 80
    protocol: TCP
```

#### readinessProbe
定义deployment等资源中的就绪探针，如：
```yaml
readinessProbe:
  exec:
    command:
      - cat
      - /tmp/healthy
  initialDelaySeconds: 5
  periodSeconds: 5
```

#### livenessProbe
定义deployment等资源中的存活探针，如：
```yaml
livenessProbe:
  httpGet:
    path: /healthz
    port: liveness-port
  failureThreshold: 1
  periodSeconds: 10
```

#### startupProbe
定义deployment等资源中的启动探针，如：
```yaml
startupProbe:
  httpGet:
    path: /healthz
    port: liveness-port
  failureThreshold: 30
  periodSeconds: 10
```

#### volumes
定义volumes，如：
```yaml
volumes:
  - name: pvc-sfs-example
    persistentVolumeClaim:
      claimName:  pvc-example
```

#### volumeMounts
定义volumeMounts，如：
```yaml
volumeMounts: 
  - mountPath: /tmp    # 挂载路径
    name: pvc-sfs-example
```
#### nodeSelector
定义nodeSelector,如：
```yaml
nodeSelector:      # 节点选择，当节点拥有gpu=true时才在节点上创建Pod
  gpu: true
```

#### affinity
定义affinity,如：
```yaml
affinity:
  nodeAffinity:
    requiredDuringSchedulingIgnoredDuringExecution:
      nodeSelectorTerms:
        - matchExpressions:
          - key: gpu
            operator: In
            values:
            - "true"
```

#### tolerations
定义tolerations，如：
```yaml
tolerations:
  - key: "key1"
    operator: "Equal"
    value: "value1"
    effect: "NoSchedule" 
```

#### hostAliases
定义hostAliases，如：
```yaml
hostAliases:
  - ip: 127.0.0.1
    hostnames:
    - foo.local
    - bar.local
  - ip: 10.1.2.3
    hostnames:
    - foo.remote
    - bar.remote
```

#### initContainers



#### podSecurityContext
定义pod的securityContext，如：
```yaml
podSecurityContext:
  enabled: true
  fsGroup: 1001
```

#### containerSecurityContext
定义容器的securityContext，如：
```yaml
containerSecurityContext:
  enabled: true
  runAsUser: 1001
  runAsNonRoot: true
```

#### hostNetwork
定义hostNetwork，如：
```yaml
hostNetwork: true
```

#### rbac 相关 （role、clusterRole、roleBinding、clusterRoleBinding）
基本的字段如下：
```yaml
rbac:
  clusterRole:
    create: false
    name: ""
    rules: ""
  role:
    create: false
    name: ""
  clusterRoleBinding:
    create: true
    name: "agentorca-crb"
    serviceAccountName: ""
    clusterroleName: "cluster-admin"
  roleBinding:
    create: false
    name: ""
    serviceAccountName: ""
    roleName: ""
```

## 改造我们自己的项目
改造我们自己的服务，有两种方式，第一种是在原先yaml格式的资源文件中使用 {{ include }} 方式，引入common中定义的如names等非deployment、service模板。
第二种是完全使用common中提供的deployment、service模板，将原先yaml格式的资源文件，全部拆解到values.yaml中。优先使用第二种方式。

我们看一个较为完整的案例，改造`openresty` 的deployment和service。
### 使用common中的deployment和service

首先原始deployment.yaml如下：

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: openresty
  namespace: cloudtogo-system
spec:
  replicas: 1
  selector:
    matchLabels:
      name: openresty
  template:
    metadata:
      labels:
        name: openresty
    spec:
      restartPolicy: Always
      containers:       
        - name: openresty
          image: "{{ .Values.global.dockerRegisty }}/{{ .Values.global.openrestyImage }}"
          #===========先以系统环境变量传进来，再用$()使用它=============
          env:  #在环境变量中使用configmap
          - name: CLOUDOS_FILESTORAGE_SERVICE_ADDRESS
            valueFrom:
              configMapKeyRef:
                name: cm-openresty
                key: CLOUDOS_FILESTORAGE_SERVICE_ADDRESS
          - name: CLOUDOS_WEBX_SERVICE_ADDRESS
            valueFrom:
              configMapKeyRef:
                name: cm-openresty
                key: CLOUDOS_WEBX_SERVICE_ADDRESS
          - name: OPENRESTY_JWT_SIGNING_KEY
            valueFrom:
              configMapKeyRef:
                name: cm-openresty
                key: OPENRESTY_JWT_SIGNING_KEY
          - name: TITAN_IDE_WORKSPACE_ADDRESS
            valueFrom:
              configMapKeyRef:
                name: cm-openresty
                key: TITAN_IDE_WORKSPACE_ADDRESS
          readinessProbe:
            tcpSocket:
              port: 80
            initialDelaySeconds: 5
            periodSeconds: 5
```

以上deployment，我们需要在values.yaml定义以下几个内容：

- 资源名称(fullnameOverride)
- 副本数(replicaCount)
- 镜像名称(image)
- deployment匹配pod的matchLabels(matchLabels)
- 环境变量(env)
- 就绪探针(readinessProbe)

然后修改values.yaml文件内容：

```yaml
fullnameOverride: "openresty"

replicaCount: 1

image:
  name: "{{ .Values.global.dockerRegisty }}/{{ .Values.global.openrestyImage }}"

matchLabels:
  name: openresty

env:
  - name: CLOUDOS_FILESTORAGE_SERVICE_ADDRESS
    valueFrom:
      configMapKeyRef:
        name: cm-openresty
        key: CLOUDOS_FILESTORAGE_SERVICE_ADDRESS
  - name: CLOUDOS_WEBX_SERVICE_ADDRESS
    valueFrom:
      configMapKeyRef:
        name: cm-openresty
        key: CLOUDOS_WEBX_SERVICE_ADDRESS
  - name: OPENRESTY_JWT_SIGNING_KEY
    valueFrom:
      configMapKeyRef:
        name: cm-openresty
        key: OPENRESTY_JWT_SIGNING_KEY
  - name: TITAN_IDE_WORKSPACE_ADDRESS
    valueFrom:
      configMapKeyRef:
        name: cm-openresty
        key: TITAN_IDE_WORKSPACE_ADDRESS

readinessProbe:
  tcpSocket:
    port: 80
  initialDelaySeconds: 5
  periodSeconds: 5

```
将deployment.yaml中原有内容清空，放置以下代码：

```gotemplate
{{- include "common.deployment" . }}
```
我们可以先测试以下，测试前，保证Chart.yaml中有以下代码：

```yaml
dependencies:
  - name: common
    version: "1.0.0"
    repository: "file://../../../common"
```

其中`repository`后面的common路径视具体情况而定。

确保Chart.yaml中有以上代码后，执行以下命令，将common依赖引入到当前charts中：

```yaml
helm dependency update “当前chart包名”
```

执行效果：
```shell
[root@infra charts]# helm dependency update openresty/
Hang tight while we grab the latest from your chart repositories...
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Deleting outdated charts
[root@infra charts]# 

```

ok，执行`helm dry-run` 去看看刚才改造的效果：

```shell
helm install openresty openresty/ -n example -f ~/mark/helm/cloud-playbook/main/cloudtogo/values.yaml --dry-run --debug
```

效果：
```yaml
---
# Source: openresty/templates/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: openresty
  namespace: "example"
  labels:
    app.kubernetes.io/name: openresty
    helm.sh/chart: openresty-0.1.0
    app.kubernetes.io/instance: openresty
    app.kubernetes.io/managed-by: Helm
spec:
  replicas: 1
  selector:
    matchLabels:
      name: openresty
  template:
    metadata:
      labels:
        name: openresty
    spec:
      containers:
        - name: openresty
          image: harbor.cloud2go.cn/cloudos/openresty:RC4.28.4-build-20210423201413
          env:
            - name: CLOUDOS_FILESTORAGE_SERVICE_ADDRESS
              valueFrom:
                configMapKeyRef:
                  key: CLOUDOS_FILESTORAGE_SERVICE_ADDRESS
                  name: cm-openresty
            - name: CLOUDOS_WEBX_SERVICE_ADDRESS
              valueFrom:
                configMapKeyRef:
                  key: CLOUDOS_WEBX_SERVICE_ADDRESS
                  name: cm-openresty
            - name: OPENRESTY_JWT_SIGNING_KEY
              valueFrom:
                configMapKeyRef:
                  key: OPENRESTY_JWT_SIGNING_KEY
                  name: cm-openresty
            - name: TITAN_IDE_WORKSPACE_ADDRESS
              valueFrom:
                configMapKeyRef:
                  key: TITAN_IDE_WORKSPACE_ADDRESS
                  name: cm-openresty
          readinessProbe:
            initialDelaySeconds: 5
            periodSeconds: 5
            tcpSocket:
              port: 80

[root@infra charts]# 


```
大家在执行的时候可能会出现一些错误，我上面是将除deployment.yaml文件之外的资源文件加了bak后缀，先忽略执行了。

改造完deployment，再改造service.yaml。
首先原始service.yaml如下：

```yaml
apiVersion: v1
kind: Service
metadata:
  name: openresty
  namespace: cloudtogo-system
  labels:
    app: openresty
spec:
  selector:
    name: openresty
  ports:
    - port: 80
      name: http-port
      protocol: TCP
      targetPort: 80

```
以上service，此时需要在values.yaml定义以下内容：
- 服务的端口(service)

修改values.yaml文件，在最后面增加以下内容：
```yaml
## 服务相关
service:
  ports:
    - port: 80
      name: http-port
      protocol: TCP
      targetPort: 80
```
然后清空service.yaml中原有的内容，放置如下代码：
```yaml
{{- include "common.service" . }}
```

更新依赖：
```shell
[root@infra charts]# helm dependency update openresty/
```
调试：
```shell
[root@infra charts]# helm install openresty openresty/ -n example -f ~/mark/helm/cloud-playbook/main/cloudtogo/values.yaml --dry-run --debug
```
效果：
```yaml
# Source: openresty/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
  name: openresty
  namespace: "example"
  labels:
    app.kubernetes.io/name: openresty
    helm.sh/chart: openresty-0.1.0
    app.kubernetes.io/instance: openresty
    app.kubernetes.io/managed-by: Helm
spec:
  selector:
    name: openresty
  ports:
    - name: http-port
      port: 80
      protocol: TCP
      targetPort: 80
```

### 在原有的deployment的基础上进行改造

改造后的效果如下：
```yaml
apiVersion: {{ include "common.capabilities.deployment.apiVersion" . }}
kind: Deployment
metadata:
  name: {{ include "common.names.fullname" . }}
  namespace: {{ .Release.Namespace | quote }}
  labels: {{- include "common.labels.standard" . | nindent 4 }}
  {{- if .Values.commonLabels }}
  {{- include "common.tplvalues.render" ( dict "value" .Values.commonLabels "context" $ ) | nindent 4 }}
  {{- end }}
  {{- if .Values.commonAnnotations }}
  annotations: {{- include "common.tplvalues.render" ( dict "value" .Values.commonAnnotations "context" $ ) | nindent 4 }}
  {{- end }}
spec:
  replicas: {{ required "A pod replicas is required!" .Values.replicaCount }}
  selector:
    {{- if .Values.matchLabels }}
    matchLabels: {{- include "common.tplvalues.render" ( dict "value" .Values.matchLabels "context" $ ) | nindent 6  }}
    {{- else }}
    matchLabels: {{- include "common.labels.matchLabels" . | nindent 6 }}
  {{- end }}
  template:
    metadata:
      {{- if .Values.matchLabels }}
      labels: {{- include "common.tplvalues.render" ( dict "value" .Values.matchLabels "context" $ ) | nindent 8  }}
      {{- else }}
      labels: {{- include "common.labels.matchLabels" . | nindent 8 }}
    {{- end }}
    spec:
      restartPolicy: Always
      containers:
        - name: openresty
          image: "{{ .Values.global.dockerRegisty }}/{{ .Values.global.openrestyImage }}"
          #===========先以系统环境变量传进来，再用$()使用它=============
          env:  #在环境变量中使用configmap
            - name: CLOUDOS_FILESTORAGE_SERVICE_ADDRESS
              valueFrom:
                configMapKeyRef:
                  name: cm-openresty
                  key: CLOUDOS_FILESTORAGE_SERVICE_ADDRESS
            - name: CLOUDOS_WEBX_SERVICE_ADDRESS
              valueFrom:
                configMapKeyRef:
                  name: cm-openresty
                  key: CLOUDOS_WEBX_SERVICE_ADDRESS
            - name: OPENRESTY_JWT_SIGNING_KEY
              valueFrom:
                configMapKeyRef:
                  name: cm-openresty
                  key: OPENRESTY_JWT_SIGNING_KEY
            - name: TITAN_IDE_WORKSPACE_ADDRESS
              valueFrom:
                configMapKeyRef:
                  name: cm-openresty
                  key: TITAN_IDE_WORKSPACE_ADDRESS
          readinessProbe:
            tcpSocket:
              port: 80
            initialDelaySeconds: 5
            periodSeconds: 5
```

service不多赘述，大同小异。








