/**
 * Jenkins build definition. This file defines the entire build pipeline.
 */
import java.net.URLEncoder
import jenkins.model.Jenkins

/**
  * PhabConnector handles all communication with phabricator if the build
  * was triggered by a phabricator run.
  */
class PhabConnector {
  def jenkinsCtx
  def URL
  def repository
  def apiToken
  def phid

  PhabConnector(jenkinsCtx, URL, repository, apiToken, phid) {
    this.jenkinsCtx = jenkinsCtx
    this.URL = URL
    this.repository = repository
    this.apiToken = apiToken
    this.phid = phid
  }

  def harborMasterUrl(method) {
    def url = "${URL}/api/${method}?api.token=${apiToken}" +
            "&buildTargetPHID=${phid}"
    return url
  }

  def sendBuildStatus(build_status) {
    def url = this.harborMasterUrl('harbormaster.sendmessage')
    def body = "type=${build_status}"
    jenkinsCtx.httpRequest consoleLogResponseBody: true,
      contentType: 'APPLICATION_FORM',
      httpMode: 'POST',
      requestBody: body,
      responseHandle: 'NONE',
      url: url,
      validResponseCodes: '200'
  }

  def addArtifactLink(linkURL, artifactKey, artifactName) {
    def encodedDisplayUrl = URLEncoder.encode(linkURL, 'UTF-8')
    def url = this.harborMasterUrl('harbormaster.createartifact')
    def body = ''
    body += "&buildTargetPHID=${phid}"
    body += "&artifactKey=${artifactKey}"
    body += '&artifactType=uri'
    body += "&artifactData[uri]=${encodedDisplayUrl}"
    body += "&artifactData[name]=${artifactName}"
    body += '&artifactData[ui.external]=true'

    jenkinsCtx.httpRequest consoleLogResponseBody: true,
      contentType: 'APPLICATION_FORM',
      httpMode: 'POST',
      requestBody: body,
      responseHandle: 'NONE',
      url: url,
      validResponseCodes: '200'
  }
}

/**
  * We expect the following parameters to be defined (for code review builds):
  *    PHID: Which should be the buildTargetPHID from Harbormaster.
  *    INITIATOR_PHID: Which is the PHID of the initiator (ie. Differential)
  *    API_TOKEN: The api token to use to communicate with Phabricator
  *    REVISION: The revision ID of the Differential.
  */

// NOTE: We use these without a def/type because that way Groovy will treat these as
// global variables.
phabConnector = PhabConnector.newInstance(this, 'https://phab.corp.pixielabs.ai' /*url*/,
                                          'PLM' /*repository*/, params.API_TOKEN, params.PHID)

SRC_STASH_NAME = 'src'
TARGETS_STASH_NAME = 'targets'
DEV_DOCKER_IMAGE = 'pixie-oss/pixie-dev-public/dev_image_with_extras'
DEV_DOCKER_IMAGE_EXTRAS = 'pixie-oss/pixie-dev-public/dev_image_with_extras'
GCLOUD_DOCKER_IMAGE = 'google/cloud-sdk:287.0.0'
COPYBARA_DOCKER_IMAGE = 'gcr.io/pixie-oss/pixie-dev-public/copybara:20210420'
GCS_STASH_BUCKET = 'px-jenkins-build-temp'

K8S_PROD_CLUSTER = 'https://cloud-prod.internal.corp.pixielabs.ai'
// Our staging instance used to be run on our prod cluster. These creds are
// actually the creds for our prod cluster.
K8S_PROD_CREDS = 'cloud-staging'

K8S_STAGING_CLUSTER = 'https://cloud-staging.internal.corp.pixielabs.ai'
K8S_STAGING_CREDS = 'pixie-prod-staging-cluster'

K8S_TESTING_CLUSTER = 'https://cloud-testing.internal.corp.pixielabs.ai'
K8S_TESTING_CREDS = 'pixie-prod-testing-cluster'

// PXL Docs variables.
PXL_DOCS_BINARY = '//src/carnot/docstring:docstring'
PXL_DOCS_FILE = 'pxl-docs.json'
PXL_DOCS_BUCKET = 'pl-docs'
PXL_DOCS_GCS_PATH = "gs://${PXL_DOCS_BUCKET}/${PXL_DOCS_FILE}"

// BPF Setup.
// The default kernel should be the oldest supported kernel
// to ensure that we don't have BPF compatibility regressions.
BPF_DEFAULT_KERNEL='4.14'
// A list of kernels to test. A jenkins worker with template
// named `jenkins-worker-with-${kernel}-kernel` should exist.
def BPF_KERNELS = ['4.14', '5.18.4']

def BPF_KERNELS_TO_TEST = [BPF_DEFAULT_KERNEL]

// Currently disabling TSAN on BPF builds because it runs too slow.
// In particular, the uprobe deployment takes far too long. See issue:
//    https://pixie-labs.atlassian.net/browse/PL-1329
// The benefit of TSAN on such runs is marginal anyways, because the tests
// are mostly single-threaded.
runBPFWithTSAN = false

// TODO(yzhao/oazizi): PP-2276 Fix the BPF ASAN tests.
runBPFWithASAN = false

// This variable store the dev docker image that we need to parse before running any docker steps.
devDockerImageWithTag = ''
devDockerImageExtrasWithTag = ''

stashList = []

// Flag controlling if coverage job is enabled.
isMainCodeReviewRun =  (env.JOB_NAME == 'pixie-dev/main-phab-test' || env.JOB_NAME == 'pixie-oss/build-and-test-pr')

isMainRun =  (env.JOB_NAME == 'pixie-main/build-and-test-all')
isNightlyTestRegressionRun = (env.JOB_NAME == 'pixie-main/nightly-test-regression')
isNightlyBPFTestRegressionRun = (env.JOB_NAME == 'pixie-main/nightly-test-regression-bpf')

isCopybaraPublic = env.JOB_NAME.startsWith('pixie-main/copybara-public')
isCopybaraTags = env.JOB_NAME.startsWith('pixie-main/copybara-tags')
isCopybaraPxAPI = env.JOB_NAME.startsWith('pixie-main/copybara-pxapi-go')

isOSSMainRun = (env.JOB_NAME == 'pixie-oss/build-and-test-all')
isOSSCloudBuildRun = env.JOB_NAME.startsWith('pixie-oss/cloud/')

isCLIBuildRun =  env.JOB_NAME.startsWith('pixie-release/cli/')
isOperatorBuildRun = env.JOB_NAME.startsWith('pixie-release/operator/')
isVizierBuildRun = env.JOB_NAME.startsWith('pixie-release/vizier/')
isCloudProdBuildRun = env.JOB_NAME.startsWith('pixie-release/cloud/')
isCloudStagingBuildRun = env.JOB_NAME.startsWith('pixie-release/cloud-staging/')
isStirlingPerfEval = (env.JOB_NAME == 'pixie-main/stirling-perf-eval')

// Build tags are used to modify the behavior of the build.
// Note: Tags only work for code-review builds.
// Enable the BPF build, even if it's not required.
buildTagBPFBuild = false
// Enable BPF build across all tested kernels.
buildTagBPFBuildAllKernels = false

def WithGCloud(Closure body) {
  if (env.KUBERNETES_SERVICE_HOST) {
    container('gcloud') {
      body()
    }
  } else {
    docker.image(GCLOUD_DOCKER_IMAGE).inside {
      body()
    }
  }
}

def gsutilCopy(String src, String dest) {
  WithGCloud {
    sh """
    gsutil -o GSUtil:parallel_composite_upload_threshold=150M cp ${src} ${dest}
    """
  }
}

def bbLinks() {
    def linkURL = "--build_metadata=BUILDBUDDY_LINKS='[Jenkins](${BUILD_URL})"
    if (isPhabricatorTriggeredBuild()) {
      def phabricator_link = ''
      if (params.REVISION) {
        phabricator_link = "${phabConnector.URL}/D${REVISION}"
      } else {
        phabricator_link = "${phabConnector.URL}/r${phabConnector.repository}${env.PHAB_COMMIT}"
      }
      linkURL += ",[Phabricator](${phabricator_link})"
    }
    linkURL += "'"
    return linkURL
}

def stashOnGCS(String name, String pattern, String excludes = '') {
  def extraExcludes = ''
  if (excludes.length() != 0) {
    extraExcludes = '--exclude=${excludes}'
  }

  def destFile = "${name}.tar.gz"
  sh """
    mkdir -p .archive && tar --exclude=.archive ${extraExcludes} -czf .archive/${destFile} ${pattern}
  """

  gsutilCopy(".archive/${destFile}", "gs://${GCS_STASH_BUCKET}/${env.BUILD_TAG}/${destFile}")
}

def unstashFromGCS(String name) {
  def srcFile = "${name}.tar.gz"
  sh 'mkdir -p .archive'

  gsutilCopy("gs://${GCS_STASH_BUCKET}/${env.BUILD_TAG}/${srcFile}", ".archive/${srcFile}")

  // Note: The tar extraction must use `--no-same-owner`.
  // Without this, the owner of some third_party files become invalid users,
  // which causes some cmake projects to fail with "failed to preserve ownership" messages.
  sh """
    tar -zxf .archive/${srcFile} --no-same-owner
    rm -f .archive/${srcFile}
  """
}

def shFileExists(String f) {
  return sh(
    script: "test -f ${f}",
    returnStatus: true) == 0
}

def shFileEmpty(String f) {
  return sh(
    script: "test -s ${f}",
    returnStatus: true) != 0
}
/**
  * @brief Add build info to harbormaster and badge to Jenkins.
  */
def addBuildInfo = {
  phabConnector.addArtifactLink(env.RUN_DISPLAY_URL, 'jenkins.uri', 'Jenkins')

  def text = ''
  def link = ''
  // Either a revision of a commit to main.
  if (params.REVISION) {
    def revisionId = "D${REVISION}"
    text = revisionId
    link = "${phabConnector.URL}/${revisionId}"
  } else {
    text = params.PHAB_COMMIT.substring(0, 7)
    link = "${phabConnector.URL}/r${phabConnector.repository}${env.PHAB_COMMIT}"
  }
  addShortText(
    text: text,
    background: 'transparent',
    border: 0,
    borderColor: 'transparent',
    color: '#1FBAD6',
    link: link
  )
}

/**
 * @brief Returns true if it's a phabricator triggered build.
 *  This could either be code review build or main commit.
 */
def isPhabricatorTriggeredBuild() {
  return params.PHID != null && params.PHID != ''
}

def codeReviewPreBuild = {
  phabConnector.sendBuildStatus('work')
  addBuildInfo()
}

def codeReviewPostBuild = {
  if (currentBuild.result == 'SUCCESS' || currentBuild.result == null) {
    phabConnector.sendBuildStatus('pass')
  } else {
    phabConnector.sendBuildStatus('fail')
  }
  phabConnector.addArtifactLink(env.BUILD_URL + '/doxygen', 'doxygen.uri', 'Doxygen')
}

def writeBazelRCFile() {
  sh 'cp ci/jenkins.bazelrc jenkins.bazelrc'
  if (!isMainRun) {
    // Don't upload to remote cache if this is not running main.
    sh '''
    echo "build --remote_upload_local_results=false" >> jenkins.bazelrc
    echo "build --build_metadata=ROLE=DEV" >> jenkins.bazelrc
    '''
  } else {
    // Only set ROLE=CI if this is running on main. This controls the whether this
    // run contributes to the test matrix at https://bb.corp.pixielabs.ai/tests/
    sh '''
    echo "build --build_metadata=ROLE=CI" >> jenkins.bazelrc
    '''
  }
  withCredentials([
    string(
      credentialsId: 'buildbuddy-api-key',
      variable: 'BUILDBUDDY_API_KEY'
    ),
    string(
      credentialsId: 'github-license-ratelimit',
      variable: 'GH_API_KEY'
    )
  ]) {
    def bbAPIArg = '--remote_header=x-buildbuddy-api-key=${BUILDBUDDY_API_KEY}'
    sh "echo \"build ${bbAPIArg}\" >> jenkins.bazelrc"

    def ghAPIEnv = '--action_env=GH_API_KEY=${GH_API_KEY}'
    sh "echo \"build ${ghAPIEnv}\" >> jenkins.bazelrc"
  }
}

def createBazelStash(String stashName) {
  sh 'rm -rf bazel-testlogs-archive'
  sh 'mkdir -p bazel-testlogs-archive'
  sh 'cp -a bazel-testlogs/ bazel-testlogs-archive || true'
  stashOnGCS(stashName, 'bazel-testlogs-archive/**')
  stashList.add(stashName)
}

def RetryOnK8sDownscale(Closure body, int times=5) {
  for (int retryCount = 0; retryCount < times; retryCount++) {
    try {
      body()
      return
    } catch (io.fabric8.kubernetes.client.KubernetesClientException e) {
      println("Caught ${e}, assuming K8s cluster downscaled, will retry.")
      // Sleep an extra 5 seconds for each retry attempt.
      def interval = (retryCount + 1) * 5
      sleep interval
      continue
    } catch (Exception e) {
      println("Unhandled ${e}, assuming fatal error.")
      throw e
    }
  }
}

def WithSourceCodeK8s(String suffix="${UUID.randomUUID()}", Integer timeoutMinutes=90, Closure body) {
  warnError('Script failed') {
    DefaultBuildPodTemplate(suffix) {
      timeout(time: timeoutMinutes, unit: 'MINUTES') {
        container('pxbuild') {
          sh '''
            git config --global --add safe.directory `pwd`
          '''
        }
        container('gcloud') {
          unstashFromGCS(SRC_STASH_NAME)
          sh 'cp ci/bes-k8s.bazelrc bes.bazelrc'
        }
        body()
      }
    }
  }
}

def WithSourceCodeAndTargetsK8s(String suffix="${UUID.randomUUID()}", Closure body) {
  WithSourceCodeK8s(suffix) {
    container('gcloud') {
      unstashFromGCS(TARGETS_STASH_NAME)
    }
    body()
  }
}

def WithSourceCodeAndTargetsBPFEnv(String stashName = SRC_STASH_NAME, String kernel = BPF_DEFAULT_KERNEL, Closure body) {
  warnError('Script failed') {
    WithSourceCodeFatalErrorBPFEnv(stashName, kernel, {
      unstashFromGCS(TARGETS_STASH_NAME)
      body()
    })
  }
}

/**
  * This function checks out the source code and wraps the builds steps.
  */
def WithSourceCodeFatalErrorBPFEnv(String stashName = SRC_STASH_NAME, String kernel, Closure body) {
  timeout(time: 90, unit: 'MINUTES') {
    node("jenkins-worker-with-${kernel}-kernel") {
      sh 'hostname'
      deleteDir()
      unstashFromGCS(stashName)
      sh 'cp ci/bes-gce.bazelrc bes.bazelrc'
      body()
    }
  }
}

/**
  * Our default docker step :
  *   3. Starts docker container.
  *   4. Runs the passed in body.
  */
def dockerStep(String dockerConfig = '', String dockerImage = devDockerImageWithTag, Closure body) {
  docker.withRegistry('https://gcr.io', 'gcr:pl-dev-infra') {
    jenkinsMnt = ' -v /mnt/jenkins/sharedDir:/mnt/jenkins/sharedDir'
    dockerSock = ' -v /var/run/docker.sock:/var/run/docker.sock -v /var/lib/docker:/var/lib/docker'
    // TODO(zasgar): We should be able to run this in isolated networks. We need --net=host
    // because dockertest needs to be able to access sibling containers.
    docker.image(dockerImage).inside(dockerConfig + dockerSock + jenkinsMnt + ' --net=host') {
      body()
    }
  }
}

def runBazelCmd(String f, String targetConfig, int retries = 5) {
  def retval = sh(
    script: "bazel ${f} ${bbLinks()} --build_metadata=CONFIG=${targetConfig}",
    returnStatus: true
  )

  if (retval == 38 && (targetConfig == 'tsan' || targetConfig == 'asan')) {
    // If bes update failed for a sanitizer run, re-run to get the real retval.
    if (retries == 0) {
      println('Bazel bes update failed for sanitizer run after multiple retries.')
      return retval
    }
    println('Bazel bes update failed for sanitizer run, retrying...')
    return runBazelCmd(f, targetConfig, retries - 1)
  }
  // 4 means that tests not present.
  // 38 means that bes update failed/
  // Both are not fatal.
  if (retval == 0 || retval == 4 || retval == 38) {
    if (retval != 0) {
      println("Bazel returned ${retval}, ignoring...")
    }
    return 0
  }
  return retval
}
/**
  * Runs bazel CI mode for main/phab builds.
  *
  * The targetFilter can either be a bazel filter clause, or bazel path (//..., etc.), but not a list of paths.
  */
def bazelCICmd(String name, String targetConfig='clang', String targetCompilationMode='opt',
               String targetsSuffix, String bazelRunExtraArgs='') {
  def buildableFile = "bazel_buildables_${targetsSuffix}"
  def testFile = "bazel_tests_${targetsSuffix}"

  def args = "-c ${targetCompilationMode} --config=${targetConfig} --build_metadata=COMMIT_SHA=\$(git rev-parse HEAD) ${bazelRunExtraArgs}"

  if (runBazelCmd("build ${args} --target_pattern_file ${buildableFile}", targetConfig) != 0) {
    throw new Exception('Bazel run failed')
  }
  if (runBazelCmd("test ${args} --target_pattern_file ${testFile}", targetConfig) != 0) {
    throw new Exception('Bazel test failed')
  }
  createBazelStash("${name}-testlogs")
}

def processBazelLogs(String logBase) {
  step([
    $class: 'XUnitPublisher',
    thresholds: [
      [
        $class: 'FailedThreshold',
        unstableThreshold: '1'
      ]
    ],
    tools: [
      [
        $class: 'GoogleTestType',
        skipNoTestFiles: true,
        pattern: "${logBase}/bazel-testlogs-archive/**/*.xml"
      ]
    ]
  ])
}

def processAllExtractedBazelLogs() {
  stashList.each({ stashName ->
    if (stashName.endsWith('testlogs')) {
      processBazelLogs(stashName)
    }
  })
}

def publishDoxygenDocs() {
  publishHTML([
    allowMissing: true,
    alwaysLinkToLastBuild: true,
    keepAll: true,
    reportDir: 'doxygen-docs/docs/html',
    reportFiles: 'index.html',
    reportName: 'doxygen'
  ])
}

def sendSlackNotification() {
  if (currentBuild.result != 'SUCCESS' && currentBuild.result != null) {
    slackSend color: '#FF0000', message: "FAILED: Build - ${env.BUILD_TAG} -- URL: ${env.BUILD_URL}."
  }
  else if (currentBuild.getPreviousBuild() &&
           currentBuild.getPreviousBuild().getResult().toString() != 'SUCCESS') {
    slackSend color: '#00FF00', message: "PASSED(Recovered): Build - ${env.BUILD_TAG} -- URL: ${env.BUILD_URL}."
  }
}

def sendCloudReleaseSlackNotification(String profile) {
  if (currentBuild.result == 'SUCCESS' || currentBuild.result == null) {
    slackSend color: '#00FF00', message: "${profile} Cloud deployed - ${env.BUILD_TAG} -- URL: ${env.BUILD_URL}."
  } else {
    slackSend color: '#FF0000', message: "${profile} Cloud deployed FAILED - ${env.BUILD_TAG} -- URL: ${env.BUILD_URL}."
  }
}

def postBuildActions = {
  if (isPhabricatorTriggeredBuild()) {
    codeReviewPostBuild()
  }

  // Main runs are triggered by Phabricator, but we still want
  // notifications on failure.
  if (!isPhabricatorTriggeredBuild() || isMainRun) {
    sendSlackNotification()
  }
}

def InitializeRepoState(String stashName = SRC_STASH_NAME) {
  sh './ci/save_version_info.sh'
  sh './ci/save_diff_info.sh'

  writeBazelRCFile()

  // Get docker image tag.
  def properties = readProperties file: 'docker.properties'
  devDockerImageWithTag = DEV_DOCKER_IMAGE + ":${properties.DOCKER_IMAGE_TAG}"
  devDockerImageExtrasWithTag = DEV_DOCKER_IMAGE_EXTRAS + ":${properties.DOCKER_IMAGE_TAG}"

  stashOnGCS(SRC_STASH_NAME, '.')
}

def DefaultGitPodTemplate(String suffix, Closure body) {
  RetryOnK8sDownscale {
    def label = "worker-git-${env.BUILD_TAG}-${suffix}"
    podTemplate(label: label, cloud: 'devinfra-cluster-usw1-0', containers: [
      containerTemplate(name: 'git', image: 'bitnami/git:2.33.0', command: 'cat', ttyEnabled: true)
    ]) {
      node(label) {
        body()
      }
    }
  }
}

def DefaultGCloudPodTemplate(String suffix, Closure body) {
  RetryOnK8sDownscale {
    def label = "worker-gcloud-${env.BUILD_TAG}-${suffix}"
    podTemplate(label: label, cloud: 'devinfra-cluster-usw1-0', containers: [
      containerTemplate(name: 'gcloud', image: GCLOUD_DOCKER_IMAGE, command: 'cat', ttyEnabled: true)
    ]) {
      node(label) {
        body()
      }
    }
  }
}

def DefaultCopybaraPodTemplate(String suffix, Closure body) {
  RetryOnK8sDownscale {
    def label = "worker-copybara-${env.BUILD_TAG}-${suffix}"
    podTemplate(label: label, cloud: 'devinfra-cluster-usw1-0', containers: [
      containerTemplate(name: 'copybara', image: COPYBARA_DOCKER_IMAGE, command: 'cat', ttyEnabled: true),
    ]) {
      node(label) {
        body()
      }
    }
  }
}

def DefaultBuildPodTemplate(String suffix, Closure body) {
  RetryOnK8sDownscale {
    def label = "worker-${env.BUILD_TAG}-${suffix}"
    podTemplate(
      label: label, cloud: 'devinfra-cluster-usw1-0', containers: [
        containerTemplate(
          name: 'pxbuild', image: 'gcr.io/' + devDockerImageWithTag,
          command: 'cat', ttyEnabled: true,
          resourceRequestMemory: '58368Mi',
          resourceRequestCpu: '30000m',
        ),
        containerTemplate(name: 'gcloud', image: GCLOUD_DOCKER_IMAGE, command: 'cat', ttyEnabled: true),
      ],
      yaml:'''
spec:
  dnsPolicy: ClusterFirstWithHostNet
  containers:
    - name: pxbuild
      securityContext:
        capabilities:
          add:
            - SYS_PTRACE
''',
      yamlMergeStrategy: merge(),
      hostNetwork: true,
      volumes: [
        hostPathVolume(mountPath: '/var/run/docker.sock', hostPath: '/var/run/docker.sock'),
        hostPathVolume(mountPath: '/var/lib/docker', hostPath: '/var/lib/docker'),
        hostPathVolume(mountPath: '/mnt/jenkins/sharedDir', hostPath: '/mnt/jenkins/sharedDir')
      ]) {
      node(label) {
        body()
      }
    }
  }
}

/**
 * Checkout the source code, record git info and stash sources.
 */
def checkoutAndInitialize() {
  DefaultGCloudPodTemplate('init') {
    container('gcloud') {
      deleteDir()
      checkout scm
      InitializeRepoState()
      if(isPhabricatorTriggeredBuild()) {
        def logMessage = sh (
          script: "git log origin/main..",
          returnStdout: true,
        ).trim()

        def hasTag = {log, tag -> (log ==~ "(?s).*#ci:${tag}(\\s|\$).*")}

        buildTagBPFBuild = hasTag(logMessage, 'bpf-build')
        buildTagBPFBuildAllKernels = hasTag(logMessage, 'bpf-build-all-kernels')
      }
    }
  }
}

def enableForTargets(String targetName, Closure body) {
  if (!shFileEmpty("bazel_buildables_${targetName}") || !shFileEmpty("bazel_tests_${targetName}")) {
    body()
  }
}

/*****************************************************************************
 * BUILDERS: This sections defines all the build steps that will happen in parallel.
 *****************************************************************************/
def preBuild = [:]
def builders = [:]

def buildAndTestOptWithUI = {
  WithSourceCodeAndTargetsK8s('build-opt') {
    container('pxbuild') {
      withCredentials([
        file(
          credentialsId: 'pl-dev-infra-jenkins-sa-json',
          variable: 'GOOGLE_APPLICATION_CREDENTIALS')
      ]) {
        bazelCICmd('build-opt', 'clang', 'opt', 'clang_opt', '--action_env=GOOGLE_APPLICATION_CREDENTIALS')
      }
    }
  }
}

def buildClangTidy = {
  WithSourceCodeK8s('clang-tidy') {
    container('pxbuild') {
      def stashName = 'build-clang-tidy-logs'
      if (isMainRun) {
        // For main builds we run clang tidy on changes files in the past 10 revisions,
        // this gives us a good balance of speed and coverage.
        sh 'ci/run_clang_tidy.sh -f diff_head_cc'
      } else {
        // For code review builds only run on diff.
        sh 'ci/run_clang_tidy.sh -f diff_origin_main_cc'
      }
      stashOnGCS(stashName, 'clang_tidy.log')
      stashList.add(stashName)
    }
  }
}

def buildDbg = {
  WithSourceCodeAndTargetsK8s('build-dbg') {
    container('pxbuild') {
      bazelCICmd('build-dbg', 'clang', 'dbg', 'clang_dbg', '--action_env=GOOGLE_APPLICATION_CREDENTIALS')
    }
  }
}

def buildGoRace = {
  WithSourceCodeAndTargetsK8s('build-go-race') {
    container('pxbuild') {
      bazelCICmd('build-go-race', 'go_race', 'opt', 'go_race')
    }
  }
}

def buildASAN = {
  WithSourceCodeAndTargetsK8s('build-san') {
    container('pxbuild') {
      bazelCICmd('build-asan', 'asan', 'dbg', 'sanitizer')
    }
  }
}

def buildTSAN = {
  WithSourceCodeAndTargetsK8s('build-san') {
    container('pxbuild') {
      bazelCICmd('build-tsan', 'tsan', 'dbg', 'sanitizer')
    }
  }
}

def buildGCC = {
  WithSourceCodeAndTargetsK8s('build-gcc-opt') {
    container('pxbuild') {
      bazelCICmd('build-gcc-opt', 'gcc', 'opt', 'gcc_opt')
    }
  }
}

def dockerArgsForBPFTest = '--privileged --pid=host -v /:/host -v /sys:/sys --env PL_HOST_PATH=/host'

def buildAndTestBPFOpt = { kernel ->
  WithSourceCodeAndTargetsBPFEnv(SRC_STASH_NAME, kernel, {
    dockerStep(dockerArgsForBPFTest, {
      bazelCICmd('build-bpf', 'bpf', 'opt', 'bpf')
    })
  })
}

def buildAndTestBPFASAN = { kernel ->
  WithSourceCodeAndTargetsBPFEnv(SRC_STASH_NAME, kernel, {
    dockerStep(dockerArgsForBPFTest, {
      bazelCICmd('build-bpf-asan', 'bpf_asan', 'dbg', 'bpf_sanitizer')
    })
  })
}

def buildAndTestBPFTSAN = { kernel ->
  WithSourceCodeAndTargetsBPFEnv(SRC_STASH_NAME, kernel, {
    dockerStep(dockerArgsForBPFTest, {
      bazelCICmd('build-bpf-tsan', 'bpf_tsan', 'dbg', 'bpf_sanitizer')
    })
  })
}

def generateTestTargets = {
  enableForTargets('clang_opt') {
    builders['Build & Test (clang:opt + UI)'] = buildAndTestOptWithUI
  }

//  enableForTargets('clang_tidy') {
//    builders['Clang-Tidy'] = buildClangTidy
//  }

  enableForTargets('clang_dbg') {
    builders['Build & Test (dbg)'] = buildDbg
  }

  enableForTargets('sanitizer') {
    builders['Build & Test (asan)'] = buildASAN
  }

  enableForTargets('sanitizer') {
    builders['Build & Test (tsan)'] = buildTSAN
  }

  enableForTargets('gcc_opt') {
    builders['Build & Test (gcc:opt)'] = buildGCC
  }

  enableForTargets('go_race') {
    builders['Build & Test (go race detector)'] = buildGoRace
  }

  BPF_KERNELS_TO_TEST.each { kernel ->
    enableForTargets('bpf') {
      builders["Build & Test (bpf tests - opt) - ${kernel}"] = { buildAndTestBPFOpt(kernel) }
    }

    if (runBPFWithASAN) {
      enableForTargets('bpf_sanitizer') {
        builders["Build & Test (bpf tests - asan) - ${kernel}"] = { buildAndTestBPFASAN(kernel) }
      }
    }

    if (runBPFWithTSAN) {
      enableForTargets('bpf_sanitizer') {
        builders["Build & Test (bpf tests - tsan) - ${kernel}"] = { buildAndTestBPFTSAN(kernel) }
      }
    }
  }
}

preBuild['Process Dependencies'] = {
  WithSourceCodeK8s('process-deps') {
    container('pxbuild') {
      def forceAll = ''
      def enableBPF = ''

      if (isMainRun || isNightlyTestRegressionRun || isOSSMainRun || isNightlyBPFTestRegressionRun) {
        forceAll = '-a'
        enableBPF = '-b'
      }

      if (buildTagBPFBuild || buildTagBPFBuildAllKernels) {
        enableBPF = '-b'
      }

      sh """
      ./ci/bazel_build_deps.sh ${forceAll} ${enableBPF}
      wc -l bazel_*
      """

      if (buildTagBPFBuildAllKernels) {
        BPF_KERNELS_TO_TEST = BPF_KERNELS
      }

      stashOnGCS(TARGETS_STASH_NAME, 'bazel_*')
      generateTestTargets()
    }
  }
}

if (isMainRun || isOSSMainRun) {
  def codecovToken = 'pixie-codecov-token'
  def slug = 'pixie-labs/pixielabs'
  if (isOSSMainRun) {
    codecovToken = 'pixie-oss-codecov-token'
    slug = 'pixie-io/pixie'
  }
  // Only run coverage on main runs.
  builders['Build & Test (gcc:coverage)'] = {
    WithSourceCodeAndTargetsK8s('coverage') {
      container('pxbuild') {
        warnError('Coverage command failed') {
          withCredentials([
            string(
              credentialsId: codecovToken,
              variable: 'CODECOV_TOKEN'
            )
          ]) {
            sh "ci/collect_coverage.sh -u -t ${CODECOV_TOKEN} -b main -c `cat GIT_COMMIT` -r " + slug
          }
        }
        createBazelStash('build-gcc-coverage-testlogs')
      }
    }
  }
}

def buildScriptForOSSCloudRelease = {
  try {
    stage('Checkout code') {
      checkoutAndInitialize()
    }
    stage('Build & Push Artifacts') {
      WithSourceCodeK8s {
        container('pxbuild') {
          sh './ci/cloud_build_release.sh -p'
        }
      }
    }
  }
  catch (err) {
    currentBuild.result = 'FAILURE'
    echo "Exception thrown:\n ${err}"
    echo 'Stacktrace:'
    err.printStackTrace()
  }
  postBuildActions()
}

if (isMainRun) {
  // Only run LSIF on main runs.
  builders['LSIF (sourcegraph)'] = {
    WithSourceCodeAndTargetsK8s('lsif') {
      container('pxbuild') {
        warnError('LSIF command failed') {
          withCredentials([
            string(
              credentialsId: 'sourcegraph-api-token',
              variable: 'SOURCEGRAPH_TOKEN'
            )
          ]) {
            sh 'ci/collect_and_upload_lsif.sh -t ${SOURCEGRAPH_TOKEN} -c `cat GIT_COMMIT`'
          }
        }
      }
    }
  }

  // Only run FOSSA on main runs.
  builders['FOSSA'] = {
    WithSourceCodeAndTargetsK8s('fossa') {
      container('pxbuild') {
        warnError('FOSSA command failed') {
          withCredentials([
            string(
              credentialsId: 'fossa-api-key',
              variable: 'FOSSA_API_KEY'
            )
          ]) {
            sh 'fossa analyze --branch main'
          }
        }
      }
    }
  }
}

builders['Lint & Docs'] = {
  WithSourceCodeAndTargetsK8s('lint') {
    container('pxbuild') {
      // Prototool relies on having a main branch in this checkout, so create one tracking origin/main
      sh 'git branch main --track origin/main'
      sh 'arc lint --trace'
    }

    if (shFileExists('run_doxygen')) {
      def stashName = 'doxygen-docs'
      container('pxbuild') {
        sh 'doxygen'
      }
      container('gcloud') {
        stashOnGCS(stashName, 'docs/html')
        stashList.add(stashName)
      }
    }
  }
}

/*****************************************************************************
 * END BUILDERS
 *****************************************************************************/

def archiveBuildArtifacts = {
  DefaultGCloudPodTemplate('archive') {
    container('gcloud') {
      // Unstash the build artifacts.
      stashList.each({ stashName ->
        dir(stashName) {
          unstashFromGCS(stashName)
        }
      })

      // Remove the tests attempts directory because it
      // causes the test publisher to mark as failed.
      sh 'find . -name test_attempts -type d -exec rm -rf {} +'

      publishDoxygenDocs()

      // Archive clang-tidy logs.
      //archiveArtifacts artifacts: 'build-clang-tidy-logs/**', fingerprint: true

      // Actually process the bazel logs to look for test failures.
      processAllExtractedBazelLogs()
    }
  }
}

/********************************************
 * The build script starts here.
 ********************************************/
def buildScriptForCommits = {
  DefaultGCloudPodTemplate('root') {
    if (isMainRun || isOSSMainRun) {
      def namePrefix = 'pixie-main'
      if (isOSSMainRun) {
        namePrefix = 'pixie-oss'
      }
      // If there is a later build queued up, we want to stop the current build so
      // we can execute the later build instead.
      def q = Jenkins.get().getQueue()
      abortBuild = false
      q.getItems().each {
        if (it.task.getDisplayName() == 'build-and-test-all') {
          // Use fullDisplayName to distinguish between pixie-oss and pixie-main builds.
          if (it.task.getFullDisplayName().startsWith(namePrefix)) {
            abortBuild = true
          }
        }
      }

      if (abortBuild) {
        echo 'Stopping current build because a later build is already enqueued'
        return
      }
    }

    if (isPhabricatorTriggeredBuild()) {
      codeReviewPreBuild()
    }

    try {
      stage('Checkout code') {
        checkoutAndInitialize()
      }
      stage('Pre-Build') {
        parallel(preBuild)
      }
      stage('Build Steps') {
        parallel(builders)
      }
      stage('Archive') {
        archiveBuildArtifacts()
      }
    }
    catch (err) {
      currentBuild.result = 'FAILURE'
      echo "Exception thrown:\n ${err}"
      echo 'Stacktrace:'
      err.printStackTrace()
    }

    postBuildActions()
  }
}

/*****************************************************************************
 * REGRESSION_BUILDERS: This sections defines all the test regressions steps
 * that will happen in parallel.
 *****************************************************************************/
def BPFRegressionBuilders = [:]

BPF_KERNELS.each { kernel ->
  BPFRegressionBuilders["Test (opt) ${kernel}"] = {
    WithSourceCodeAndTargetsBPFEnv(SRC_STASH_NAME, kernel, {
      dockerStep(dockerArgsForBPFTest, {
        bazelCICmd('build-bpf', 'bpf', 'opt', 'bpf')
      })
    })
  }
}


/*****************************************************************************
 * REGRESSION_BUILDERS: This sections defines all the test regressions steps
 * that will happen in parallel.
 *****************************************************************************/
def regressionBuilders = [:]

TEST_ITERATIONS = 5

regressionBuilders['Test (opt)'] = {
  WithSourceCodeAndTargetsK8s {
    container('pxbuild') {
      runBazelCmd("test -c opt --runs_per_test ${TEST_ITERATIONS} \
        --target_pattern_file bazel_tests_clang_opt", 'opt', 1)
      createBazelStash('build-opt-testlogs')
    }
  }
}

regressionBuilders['Test (ASAN)'] = {
  WithSourceCodeAndTargetsK8s {
    container('pxbuild') {
      runBazelCmd("test --config asan --runs_per_test ${TEST_ITERATIONS} \
        --target_pattern_file bazel_tests_sanitizer", 'asan', 1)
      createBazelStash('build-asan-testlogs')
    }
  }
}

regressionBuilders['Test (TSAN)'] = {
  WithSourceCodeAndTargetsK8s {
    container('pxbuild') {
      runBazelCmd("test --config tsan --runs_per_test ${TEST_ITERATIONS} \
        --target_pattern_file bazel_tests_sanitizer", 'tsan', 1)
      createBazelStash('build-tsan-testlogs')
    }
  }
}

/*****************************************************************************
 * END REGRESSION_BUILDERS
 *****************************************************************************/

/*****************************************************************************
 * STIRLING PERF BUILDERS: Create & deploy, wait, then measure CPU use.
 *****************************************************************************/

def clusterNames = (params.CLUSTER_NAMES != null) ? params.CLUSTER_NAMES.split(',') : ['']
int numPerfEvals = (params.NUM_EVAL_RUNS != null) ? Integer.parseInt(params.NUM_EVAL_RUNS) : 5
int warmupMinutes = (params.WARMUP_MINUTES != null) ? Integer.parseInt(params.WARMUP_MINUTES) : 30
int evalMinutes = (params.EVAL_MINUTES != null) ? Integer.parseInt(params.EVAL_MINUTES) : 60
int profilerMinutes = (params.PROFILER_MINUTES != null) ? Integer.parseInt(params.PROFILER_MINUTES) : 5
int cleanupClusters = (params.CLEANUP_CLUSTERS != null) ? Integer.parseInt(params.CLEANUP_CLUSTERS) : 1
String groupName = (params.GROUP_NAME != null) ? params.GROUP_NAME : 'none'
String machineType = (params.MACHINE_TYPE != null) ? params.MACHINE_TYPE : 'n2-standard-4'
String experimentTag = (params.EXPERIMENT_TAG != null) ? params.EXPERIMENT_TAG : 'none'
String gitHashForPerfEval = (params.GIT_HASH_FOR_PERF_EVAL != null) ? params.GIT_HASH_FOR_PERF_EVAL : 'HEAD'
String imageTagForPerfEval = 'none'

def stirlingPerfBuilders = [:]

String getClusterNameDateString() {
  date = new Date()
  return date.format('yyyy-MM-dd--HHmm-ss')
}

useCluster = { String clusterName ->
  sh 'hostname'
  sh 'gcloud --version'
  sh "gcloud container clusters get-credentials ${clusterName} --project pl-pixies --zone us-west1-a"
}

deleteCluster = { String clusterName ->
  // We use 'delete || true' so that failure does not cause the entire pipeline to fail or go unstable.
  // In particular, deleteCluster is invoked when createCluster fails; this has two scenarios:
  // 1. The cluster was created, but the gcloud command failed anyway.
  // 2. The cluster was not created.
  // ... For (1) above, we expect cluster deletion to succeed.
  // ... For (2) above, we invoke deleteCluster (because it is hard to know we are in this scenario),
  // ... and we expect the command to fail, but we don't want the entire build/perf-eval to stop.
  // In general, if clusters leak through, they are eventually cleaned up by the perf eval cluster
  // cleanup cron job.
  sh "gcloud container --project pl-pixies clusters delete ${clusterName} --zone us-west1-a --quiet || true"
}

createCluster = { String clusterName ->
  retryIdx = 0
  numRetries = 3

  // We will uniquify the cluster name based on the retry count because there is some chance
  // that gcloud will refuse to create a cluster (on retry) based on a name being in use.
  // Here we create a local variable 'retryUniqueClusterName' that is distinct from 'clusterName'
  // because 'clusterName' is curried into 'oneEval'. If we clobber 'clusterName' here,
  // the currying becomes wrong and different evals will wrongly all pick up the same cluster name.
  retryUniqueClusterName = null

  createClusterScript = "scripts/create_gke_cluster.sh"
  sh 'hostname'
  sh 'gcloud components update'
  sh 'gcloud --version'
  retry(numRetries) {
    if (retryIdx > 0) {
      // Prevent leaking clusters from previous attempts.
      deleteCluster(retryUniqueClusterName)
    }
    // Uniquify the cluster name based on the retryIdx because retry attempts
    // may fail based on the pre-existing cluster name.
    retryUniqueClusterName = clusterName  + '-' + String.format('%d', retryIdx)
    sh "${createClusterScript} -S -f -n 1 -c ${retryUniqueClusterName} -m ${machineType}"
    ++retryIdx
  }
}

pxDeployForStirlingPerfEval = {
  withCredentials([
    string(
      // There are two credentials for perf-evals:
      // 1. px-staging-user-api-key: staging cloud as pixie org. member.
      // 2. px-stirling-perf-eval-user-api-key: staging cloud, as "perf-eval" (a different) org.
      // Currently using (2) above because that isolates the perf evals from updates made to staging
      // cloud by the cloud team, e.g. plugin scripts running (or not running).
      credentialsId: 'px-stirling-perf-eval-user-api-key',
      variable: 'THE_PIXIE_CLI_API_KEY'
    )
  ]) {
    withEnv(['PL_CLOUD_ADDR=staging.withpixie.dev:443']) {
      // Useful if one wants to ssh in for debugging "Jenkins only" issues.
      sh 'hostname'

      // Download the latest px binary.
      // Deploy demo apps.
      // Deploy pixie.
      sh 'curl -fsSL https://storage.googleapis.com/pixie-dev-public/cli/latest/cli_linux_amd64 -o /usr/local/bin/px'
      sh 'chmod +x /usr/local/bin/px'
      sh 'px auth login --use_api_key --api_key ${THE_PIXIE_CLI_API_KEY}'
      sh 'px demo deploy px-kafka -y -q'
      sh 'px demo deploy px-sock-shop -y -q'
      sh 'px demo deploy px-online-boutique -y -q'
      sh 'px deploy -y -q'

      // Ensure skaffold is configured with the dev. image registry.
      sh 'skaffold config set default-repo gcr.io/pl-dev-infra'
      // Regenerate the json list of artifacts targeting the images built for this eval.
      sh "skaffold build -p opt -t ${imageTagForPerfEval} --dry-run -q -f skaffold/skaffold_vizier.yaml > artifacts.json"
      // Useful for local debug, or to verify the image tags.
      sh 'cat artifacts.json'
      // Skaffold deploy using perf-eval images generated in the build & push step.
      sh 'cat artifacts.json | skaffold deploy -f skaffold/skaffold_vizier.yaml --build-artifacts -'
    }
  }
}

def pxCollectPerfInfo(String clusterName, int evalIdx, int evalMinutes, int profilerMinutes) {
  withCredentials([
    string(
      credentialsId: 'px-stirling-perf-eval-user-api-key',
      variable: 'THE_PIXIE_CLI_API_KEY'
    )
  ]) {
    withEnv(['PL_CLOUD_ADDR=staging.withpixie.dev:443']) {
      // These should have been created when un-stashing the repo info.
      assert fileExists('logs')
      assert fileExists('logs/pod_resource_usage')

      // Show the cluster name (useful if results are strange and we suspect the wrong
      // cluster was used for recording perf info).
      sh 'kubectl config current-context'

      sh 'px auth login --use_api_key --api_key ${THE_PIXIE_CLI_API_KEY}'
      sh "px run -f logs/pod_resource_usage -o json -- --start_time=-${evalMinutes}m 1> logs/perf.jsons 2> logs/perf.jsons.stderr"

      sh "px run px/perf_flamegraph -o json -- --start_time=-${profilerMinutes}m --pct_basis_entity=node --pod=pem 1> logs/stack-traces.jsons 2> logs/stack-traces.jsons.stderr"
      sh "gcloud container clusters list --project pl-pixies --filter='name:${clusterName}' --format=json | tee logs/cluster-info.json"

      // Save the original results.
      indexedEvalResultName = String.format('perf-eval-results-%02d', evalIdx)
      stashOnGCS(indexedEvalResultName, 'logs')
    }
  }
}

insertRecordsToPerfDB = { int evalIdx ->
  perf_reqs = 'src/stirling/private/scripts/perf/requirements.txt'
  perf_eval = 'src/stirling/private/scripts/perf/perf-eval.py'
  withCredentials([
    usernamePassword(
      credentialsId: 'stirling-perf-postgres',
      usernameVariable: 'STIRLING_PERF_DB_USERNAME',
      passwordVariable: 'STIRLING_PERF_DB_PASSWORD',
    )
  ]) {
    // perf-eval.py will read the git repo to find the commit hash & date/time.
    // Here, we just fail fast in case the git repo is missing.
    assert fileExists('.git')

    // Ensure git is configured correctly, and show repo state.
    sh 'git config --global --add safe.directory $(pwd)'
    sh 'git rev-parse HEAD'

    // Ensure requirements setup for perf-eval.py.
    sh "pip3 install -r ${perf_reqs}"

    // Insert the perf records into the perf db.
    // Retries exist because in rare cases, the perf db complains about too many API requests.
    numRetries = 3
    retry(numRetries) {
      sh "python3 ${perf_eval} insert-perf-records --jenkins --group-name ${groupName} --tag ${experimentTag} --idx ${evalIdx}"
    }
  }
}

def getCurrentClusterName(String clusterName) {
  def currentClusterName = sh(
    script: 'kubectl config current-context',
    returnStdout: true,
    returnStatus: false
  ).trim()

  // currentClusterName will look something like this:
  // gke_pl-pixies_us-west1-a_stirling-perf-2022-08-24--0648-09-00-0
  // However, fro a GKE perspective the name is stirling-perf-2022-08-24--0648-09-00-0.
  def tokens = currentClusterName.split('_')
  currentClusterName = tokens.last()
  return currentClusterName
}

oneEval = { int evalIdx, String clusterName, boolean newClusterNeeded ->
  int margin = 2
  int clusterCreationMinutes = 15
  int pixieDeployMinutes = 10
  int timeoutMinutes = margin * (clusterCreationMinutes + pixieDeployMinutes + warmupMinutes + evalMinutes)

  return {
    WithSourceCodeK8s('stirling-perf-eval', timeoutMinutes) {
      container('pxbuild') {

        // Unstash the "as built" repo info (see buildAndPushPemImagesForPerfEval).
        // In more detail, here, we start with a fresh fully up-to-date source tree. The "as built" repo
        // state will often be different (e.g. a particular diff or local experiment).
        // That state is only known inside of buildAndPushPemImagesForPerfEval. Because we need that
        // information, buildAndPushPemImagesForPerfEval is responsible for stashing the info on GCS,
        // and here, we recover the saved state (in file 'logs/perf_eval_repo_info.bin').
        // The stash on GCS is needed because file system state is volatile in these build stages.
        unstashFromGCS('perf-eval-repo-info')
        assert fileExists('logs/perf_eval_repo_info.bin')
        assert fileExists('logs/pod_resource_usage')

        if (newClusterNeeded) {
          // Default behavior: create a new cluster for this perf eval.
          stage("Create cluster.") {
            createCluster(clusterName)
          }
        } else {
          // A pre-existing cluster name was supplied to the build.
          stage("Use cluster.") {
            echo "clusterName: ${clusterName}."
            useCluster(clusterName)
          }
        }
        stage('Deploy pixie.') {
          pxDeployForStirlingPerfEval()
        }
        stage('Warmup.') {
          sh "sleep ${60 * warmupMinutes}"
        }
        stage('Evaluate.') {
          sh "sleep ${60 * evalMinutes}"
        }
        stage('Collect.') {
          pxCollectPerfInfo(getCurrentClusterName(clusterName), evalIdx, evalMinutes, profilerMinutes)
        }
        stage('Insert records to perf db.') {
          insertRecordsToPerfDB(evalIdx)
        }
        if (newClusterNeeded) {
          // Earlier, we had created a new cluster for this perf eval.
          // Here, we clean up.
          stage("Delete cluster.") {
            if(cleanupClusters) {
              deleteCluster(getCurrentClusterName(clusterName))
            } else {
              sh "echo skipping cluster cleanup."
            }
          }
        }
      }
    }
  }
}

def savePodResourceUsagePxlScript() {
  pod_resource_usage_path = "src/pxl_scripts/private/b7ca1b62-6c9f-4a3f-a45d-a5bdffbcae6a/pod_resource_usage"
  assert fileExists(pod_resource_usage_path)
  sh 'mkdir -p logs/pod_resource_usage'
  sh "cp ${pod_resource_usage_path}/* logs/pod_resource_usage"
}

def saveRepoInfo() {
  perf_reqs = 'src/stirling/private/scripts/perf/requirements.txt'
  perf_eval = 'src/stirling/private/scripts/perf/perf-eval.py'
  withCredentials([
    usernamePassword(
      credentialsId: 'stirling-perf-postgres',
      usernameVariable: 'STIRLING_PERF_DB_USERNAME',
      passwordVariable: 'STIRLING_PERF_DB_PASSWORD',
    )
  ]) {
    sh 'mkdir -p logs'
    sh "pip3 install -r ${perf_reqs}"
    sh "python3 ${perf_eval} save-perf-record-repo-info-to-disk --jenkins"
    stashOnGCS('perf-eval-repo-info', 'logs')
  }
}

def checkIfRequiredImagesExist() {
  numImages = Integer.parseInt(
    sh(
      script: "cat artifacts.json | jq '.builds[].imageName' | wc -l",
      returnStdout: true,
      returnStatus: false
    ).trim()
  )

  // Use the artifacts.json file and jq to build a list of all required images.
  def requiredImages = []

  for( int i=0; i < numImages; i++ ) {
    imageNameAndTag = sh(script: "cat artifacts.json | jq '.builds[${i}].tag'", returnStdout: true, returnStatus: false).trim()
    requiredImages.add(imageNameAndTag)
  }

  // allRequiredImagesExist will be set to false if we cannot find any one of the required images.
  boolean allRequiredImagesExist = true

  for (imageNameAndTag in requiredImages) {
    echo "Checking if image: ${imageNameAndTag} exists."
    describeStatusCode = sh(script: "gcloud container images describe ${imageNameAndTag}", returnStdout: false, returnStatus: true)

    if (describeStatusCode != 0) {
      echo "Image: ${imageNameAndTag} does not exist."
      allRequiredImagesExist = false
      break
    }
    else {
      echo "Image: ${imageNameAndTag} exists."
    }
  }

  if (allRequiredImagesExist) {
    sep = "\n... "
    echo "All images found:${sep}${requiredImages.join(sep)}"
  }
  return allRequiredImagesExist
}

def checkoutTargetRepo(String gitHashForPerfEval) {
  // Log out initial repo state.
  sh 'echo "Starting repo state:" && git rev-parse HEAD'

  if (params.DIFF_ID != "") {
    sshagent(['build-bot-ro']) {
      // DIFF_ID branch.
      // Specifying DIFF_ID (from Phab) enables a perf eval on an unmerged branch that resides in phab.
      // To eval this repo state, we fetch the specific tag from the staging repo & merge.
      def diffId = Integer.parseInt(params.DIFF_ID)
      sh 'mkdir -p ~/.ssh'
      sh 'ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts'
      sh 'git config remote.staging.url ssh://git@github.com/pixie-labs/pixielabs-staging.git'
      sh "git fetch --tags --force -q -- ssh://git@github.com/pixie-labs/pixielabs-staging.git refs/tags/phabricator/diff/${diffId}"
      gitHashForPerfEval = sh(script: "git rev-parse HEAD", returnStdout: true, returnStatus: false).trim()
      def targetHash = sh(script: "git rev-parse refs/tags/phabricator/diff/${diffId}^{commit}", returnStdout: true, returnStatus: false).trim()
      echo "Merging based on DIFF_ID: ${diffId}, found targetHash: ${targetHash}."
      sh "git merge --ff ${targetHash}"
      imageTagForPerfEval = 'perf-eval-' + gitHashForPerfEval + "-B${diffId}"
    }
  } else {
    // GIT_HASH_FOR_PERF_EVAL branch.
    // Here, we evaluate some commit that is merged into main.
    // Alternately (to a SHA), the user can specify a string like "HEAD~3" or "some-branch".
    // Build arg. GIT_HASH_FOR_PERF_EVAL is converted into sha,
    // and used to construct the resulting image tag.
    sh "echo 'Target repo state:' && git rev-parse ${gitHashForPerfEval}"
    gitHashForPerfEval = sh(script: "git rev-parse ${gitHashForPerfEval}", returnStdout: true, returnStatus: false).trim()
    sh "git checkout ${gitHashForPerfEval}"
    imageTagForPerfEval = 'perf-eval-' + gitHashForPerfEval
  }

  echo "Image tag for perf eval: ${imageTagForPerfEval}"
  sh 'echo "Repo state:" && git rev-parse HEAD'
  return imageTagForPerfEval
}

buildAndPushPemImagesForPerfEval = {
  WithSourceCodeK8s('pem-build-push') {
    container('pxbuild') {
      // We will need the repo, fail fast here if it is not available.
      assert fileExists('.git')

      // Ensure repo is configured for use.
      sh 'git config --global --add safe.directory $(pwd)'

      // Copy the pod resource utilization script into the logs directory,
      // so that it is stashed along with repo info.
      savePodResourceUsagePxlScript()

      imageTagForPerfEval = checkoutTargetRepo(gitHashForPerfEval)
      saveRepoInfo()

      // Ensure skaffold is configured for dev. image registry.
      sh 'skaffold config set default-repo gcr.io/pl-dev-infra'

      // Remote caching setup does not work correctly at this time:
      // disable remote caching by removing this bazelrc file.
      sh 'rm bes.bazelrc'

      // Save the image names & tags into artiacts.json, and log out the same info.
      // Useful if one wants to cross check vs. the artifacts that we deploy later.
      sh "skaffold build -p opt -t ${imageTagForPerfEval} -f skaffold/skaffold_vizier.yaml -q --dry-run | tee artifacts.json"

      allRequiredImagesExist = checkIfRequiredImagesExist()

      if (!allRequiredImagesExist) {
        echo "Building all images."
        sh "skaffold build -p opt -t ${imageTagForPerfEval} -f skaffold/skaffold_vizier.yaml"
      }
    }
  }
}

if(clusterNames[0].size()) {
  // Useful for:
  // ... debugging
  // ... faster runs or iterations
  // ... other special cases or special setups.
  // This branch allows a user to specify which cluster(s) to run the perf eval on.
  // (It will *not* create new clusters.)
  // To enable, specify the cluster name(s) as a build param. For more than one cluster,
  // use a comma separated list:
  // my-dev-cluster-00,my-dev-cluster-01
  boolean newClusterNeeded = false
  clusterNames.eachWithIndex { clusterName, i ->
    title = "Eval ${i}."
    perfEvaluator = oneEval.curry(i).curry(clusterName).curry(newClusterNeeded)
    stirlingPerfBuilders[title] = perfEvaluator()
  }
} else {
  // Default path: no cluster names supplied to the build.
  // The perf evals will create clusters.
  boolean newClusterNeeded = true
  for( int i=0; i < numPerfEvals; i++ ) {
    clusterName = 'stirling-perf-' + getClusterNameDateString() + '-' + String.format('%02d', i)
    title = "Eval ${i}."
    perfEvaluator = oneEval.curry(i).curry(clusterName).curry(newClusterNeeded)
    stirlingPerfBuilders[title] = perfEvaluator()
  }
}

/*****************************************************************************
 * END STIRLING PERF BUILDERS
 *****************************************************************************/


def buildScriptForNightlyTestRegression = { testjobs ->
  try {
    stage('Checkout code') {
      checkoutAndInitialize()
    }
    stage('Pre-Build') {
      parallel(preBuild)
    }
    stage('Testing') {
      parallel(testjobs)
    }
    stage('Archive') {
      DefaultGCloudPodTemplate('archive') {
        container('gcloud') {
          // Unstash the build artifacts.
          stashList.each({ stashName ->
            dir(stashName) {
              unstashFromGCS(stashName)
            }
          })

          // Remove the tests attempts directory because it
          // causes the test publisher to mark as failed.
          sh 'find . -name test_attempts -type d -exec rm -rf {} +'

          // Actually process the bazel logs to look for test failures.
          processAllExtractedBazelLogs()
        }
      }
    }
  }
  catch (err) {
    currentBuild.result = 'FAILURE'
    echo "Exception thrown:\n ${err}"
    echo 'Stacktrace:'
    err.printStackTrace()
  }

  postBuildActions()
}

def updateVersionsDB(String credsName, String clusterURL, String namespace) {
  WithSourceCodeK8s {
    container('pxbuild') {
      unstashFromGCS('versions')
      withKubeConfig([
        credentialsId: credsName,
        serverUrl: clusterURL,
        namespace: namespace
      ]) {
        sh './ci/update_artifact_db.sh'
      }
    }
  }
}

def  buildScriptForCLIRelease = {
  DefaultGCloudPodTemplate('root') {
    withCredentials([
      string(
        credentialsId: 'docker_access_token',
        variable: 'DOCKER_TOKEN'
      ),
      string(
        credentialsId: 'buildbot-gpg-key-id',
        variable: 'BUILDBOT_GPG_KEY_ID'
      ),
      string(
        credentialsId: 'buildbot-github-token',
        variable: 'GITHUB_TOKEN'
      )
    ]) {
      try {
        stage('Checkout code') {
          checkoutAndInitialize()
        }
        stage('Build & Push Artifacts') {
          WithSourceCodeK8s {
            container('pxbuild') {
              withCredentials([
                file(
                  credentialsId: 'buildbot-private-key-asc',
                  variable: 'BUILDBOT_GPG_KEY_FILE'
                )
              ]) {
                sh 'docker login -u pixielabs -p $DOCKER_TOKEN'
                sh './ci/cli_build_release.sh'
                stash name: 'ci_scripts_signing', includes: 'ci/**'
                stashOnGCS('versions', 'src/utils/artifacts/artifact_db_updater/VERSIONS.json')
                stashList.add('versions')
              }
            }
          }
        }
        stage('Sign Mac Binaries') {
          node('macos') {
            deleteDir()
            unstash 'ci_scripts_signing'
            withCredentials([
              file(
                credentialsId: 'buildbot-private-key-asc',
                variable: 'BUILDBOT_GPG_KEY_FILE'
              ),
              string(
                credentialsId: 'pl_ac_passwd',
                variable: 'AC_PASSWD'
              ),
              string(
                credentialsId: 'jenkins_keychain_pw',
                variable: 'JENKINSKEY'
              )
            ]) {
              sh './ci/cli_merge_sign.sh'
            }
            stash name: 'cli_darwin_signed', includes: 'cli_darwin*'
          }
        }
        stage('Upload Signed Binary') {
          node('macos') {
            WithSourceCodeK8s {
              container('pxbuild') {
                withCredentials([
                  file(
                    credentialsId: 'buildbot-private-key-asc',
                    variable: 'BUILDBOT_GPG_KEY_FILE'
                  )
                ]) {
                  unstash 'cli_darwin_signed'
                  sh './ci/cli_upload_signed.sh'
                }
              }
            }
          }
        }
        stage('Update versions database (testing)') {
          updateVersionsDB(K8S_TESTING_CREDS, K8S_TESTING_CLUSTER, 'plc-testing')
        }
        stage('Update versions database (staging)') {
          updateVersionsDB(K8S_STAGING_CREDS, K8S_STAGING_CLUSTER, 'plc-staging')
        }
        stage('Update versions database (prod)') {
          updateVersionsDB(K8S_PROD_CREDS, K8S_PROD_CLUSTER, 'plc')
        }
      }
      catch (err) {
        currentBuild.result = 'FAILURE'
        echo "Exception thrown:\n ${err}"
        echo 'Stacktrace:'
        err.printStackTrace()
      }
    }

    postBuildActions()
  }
}

def updatePxlDocs() {
  WithSourceCodeK8s {
    container('pxbuild') {
      def pxlDocsOut = "/tmp/${PXL_DOCS_FILE}"
      sh "bazel run ${PXL_DOCS_BINARY} -- --output_json ${pxlDocsOut}"
      sh "gsutil cp ${pxlDocsOut} ${PXL_DOCS_GCS_PATH}"
    }
  }
}

def vizierReleaseBuilders = [:]

vizierReleaseBuilders['Build & Push Artifacts'] = {
  WithSourceCodeK8s {
    container('pxbuild') {
      withKubeConfig([
        credentialsId: K8S_PROD_CREDS,
        serverUrl: K8S_PROD_CLUSTER, namespace: 'default'
      ]) {
        sh './ci/vizier_build_release.sh'
        stashOnGCS('versions', 'src/utils/artifacts/artifact_db_updater/VERSIONS.json')
        stashList.add('versions')
      }
    }
  }
}

vizierReleaseBuilders['Build & Export Docs'] = {
  updatePxlDocs()
}

def buildScriptForVizierRelease = {
  try {
    stage('Checkout code') {
      checkoutAndInitialize()
    }
    stage('Build & Push Artifacts') {
      parallel(vizierReleaseBuilders)
    }
    stage('Update versions database (testing)') {
      updateVersionsDB(K8S_TESTING_CREDS, K8S_TESTING_CLUSTER, 'plc-testing')
    }
    stage('Update versions database (staging)') {
      updateVersionsDB(K8S_STAGING_CREDS, K8S_STAGING_CLUSTER, 'plc-staging')
    }
    stage('Update versions database (prod)') {
      updateVersionsDB(K8S_PROD_CREDS, K8S_PROD_CLUSTER, 'plc')
    }
  }
  catch (err) {
    currentBuild.result = 'FAILURE'
    echo "Exception thrown:\n ${err}"
    echo 'Stacktrace:'
    err.printStackTrace()
  }

  postBuildActions()
}

def buildScriptForOperatorRelease = {
  try {
    stage('Checkout code') {
      checkoutAndInitialize()
    }
    stage('Build & Push Artifacts') {
      WithSourceCodeK8s {
        container('pxbuild') {
          withKubeConfig([
            credentialsId: K8S_PROD_CREDS,
            serverUrl: K8S_PROD_CLUSTER, namespace: 'default'
          ]) {
            sh './ci/operator_build_release.sh'
            stashOnGCS('versions', 'src/utils/artifacts/artifact_db_updater/VERSIONS.json')
            stashList.add('versions')
          }
        }
      }
    }
    stage('Update versions database (testing)') {
      updateVersionsDB(K8S_TESTING_CREDS, K8S_TESTING_CLUSTER, 'plc-testing')
    }
    stage('Update versions database (staging)') {
      updateVersionsDB(K8S_STAGING_CREDS, K8S_STAGING_CLUSTER, 'plc-staging')
    }
    stage('Update versions database (prod)') {
      updateVersionsDB(K8S_PROD_CREDS, K8S_PROD_CLUSTER, 'plc')
    }
  }
  catch (err) {
    currentBuild.result = 'FAILURE'
    echo "Exception thrown:\n ${err}"
    echo 'Stacktrace:'
    err.printStackTrace()
  }

  postBuildActions()
}

def pushAndDeployCloud(String profile, String namespace, String clusterCreds, String clusterURL) {
  WithSourceCodeK8s {
    container('pxbuild') {
      withKubeConfig([
        credentialsId: clusterCreds,
        serverUrl: clusterURL, namespace: namespace
      ]) {
        withCredentials([
          file(
            credentialsId: 'pl-dev-infra-jenkins-sa-json',
            variable: 'GOOGLE_APPLICATION_CREDENTIALS'
          )
        ]) {
          if (profile == 'prod') {
            sh './ci/cloud_build_release.sh -r'
          } else {
            sh './ci/cloud_build_release.sh'
          }
        }
      }
    }
  }
}

def buildScriptForCloudStagingRelease = {
  try {
    stage('Checkout code') {
      checkoutAndInitialize()
    }
    stage('Build & Push Artifacts') {
      pushAndDeployCloud('staging', 'plc-staging', K8S_STAGING_CREDS, K8S_STAGING_CLUSTER)
    }
  }
  catch (err) {
    currentBuild.result = 'FAILURE'
    echo "Exception thrown:\n ${err}"
    echo 'Stacktrace:'
    err.printStackTrace()
  }
  sendCloudReleaseSlackNotification('Staging')
  postBuildActions()
}

def buildScriptForCloudProdRelease = {
  try {
    stage('Checkout code') {
      checkoutAndInitialize()
    }
    stage('Build & Push Artifacts') {
      pushAndDeployCloud('prod', 'plc', K8S_PROD_CREDS, K8S_PROD_CLUSTER)
    }
  }
  catch (err) {
    currentBuild.result = 'FAILURE'
    echo "Exception thrown:\n ${err}"
    echo 'Stacktrace:'
    err.printStackTrace()
  }
  sendCloudReleaseSlackNotification('Prod')
  postBuildActions()
}

def copybaraTemplate(String name, String copybaraFile) {
  DefaultCopybaraPodTemplate(name) {
    deleteDir()
    checkout scm
    container('copybara') {
      sshagent (credentials: ['pixie-copybara-git']) {
        withCredentials([
          file(
            credentialsId: 'copybara-private-key-asc',
            variable: 'COPYBARA_GPG_KEY_FILE'
          ),
          string(
            credentialsId: 'copybara-gpg-key-id',
            variable: 'COPYBARA_GPG_KEY_ID'
          ),
        ]) {
          sh "GIT_SSH_COMMAND='ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no' \
          ./ci/run_copybara.sh ${copybaraFile}"
        }
      }
    }
  }
}

def buildScriptForCopybaraPublic() {
  try {
    stage('Copybara it') {
      copybaraTemplate('public-copy', 'tools/copybara/public/copy.bara.sky')
    }
    stage('Copy tags') {
      DefaultGitPodTemplate('public-copy-tags') {
        container('git') {
          deleteDir()
          checkout([
            changelog: false,
            poll: false,
            scm: [
              $class: 'GitSCM',
              branches: [[name: 'main']],
              extensions: [
                [$class: 'RelativeTargetDirectory', relativeTargetDir: 'pixie-private'],
                [$class: 'CloneOption', noTags: false, reference: '', shallow: false]
              ],
              userRemoteConfigs: [
                [credentialsId: 'build-bot-ro', url: 'git@github.com:pixie-labs/pixielabs.git']
              ]
            ]
          ])
          checkout([
            changelog: false,
            poll: false,
            scm: [
              $class: 'GitSCM',
              branches: [[name: 'main']],
              extensions: [
                [$class: 'RelativeTargetDirectory', relativeTargetDir: 'pixie-oss'],
                [$class: 'CloneOption', noTags: false, reference: '', shallow: false]
              ],
              userRemoteConfigs: [
                [credentialsId: 'pixie-copybara-git', url: 'git@github.com:pixie-io/pixie.git']
              ]
            ]
          ])
          dir('pixie-private') {
            sshagent (credentials: ['pixie-copybara-git']) {
              sh "GIT_SSH_COMMAND='ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no' \
              ./ci/copy_release_tags.sh ../pixie-oss"
            }
          }
        }
      }
    }
  }
  catch (err) {
    currentBuild.result = 'FAILURE'
    echo "Exception thrown:\n ${err}"
    echo 'Stacktrace:'
    err.printStackTrace()
  }
}

def buildScriptForCopybaraPxAPI() {
  try {
    stage('Copybara it') {
      copybaraTemplate('pxapi-copy', 'tools/copybara/pxapi_go/copy.bara.sky')
    }
  }
  catch (err) {
    currentBuild.result = 'FAILURE'
    echo "Exception thrown:\n ${err}"
    echo 'Stacktrace:'
    err.printStackTrace()
  }
}

def buildScriptForStirlingPerfEval = {
  stage('Checkout code.') {
    checkoutAndInitialize()
  }
  stage('Build & push.') {
    buildAndPushPemImagesForPerfEval()
  }
  if (currentBuild.result == 'SUCCESS' || currentBuild.result == null) {
    stage('Stirling perf eval.') {
      parallel(stirlingPerfBuilders)
    }
  }
  else {
    currentBuild.result = 'FAILURE'
  }
}

if (isNightlyTestRegressionRun) {
  buildScriptForNightlyTestRegression(regressionBuilders)
} else if (isNightlyBPFTestRegressionRun) {
  buildScriptForNightlyTestRegression(BPFRegressionBuilders)
} else if (isCLIBuildRun) {
  buildScriptForCLIRelease()
} else if (isVizierBuildRun) {
  buildScriptForVizierRelease()
} else if (isOperatorBuildRun) {
  buildScriptForOperatorRelease()
} else if (isCloudStagingBuildRun) {
  buildScriptForCloudStagingRelease()
} else if (isCloudProdBuildRun) {
  buildScriptForCloudProdRelease()
} else if (isOSSCloudBuildRun) {
  buildScriptForOSSCloudRelease()
} else if (isCopybaraPublic || isCopybaraTags) {
  buildScriptForCopybaraPublic()
} else if (isCopybaraPxAPI) {
  buildScriptForCopybaraPxAPI()
} else if (isStirlingPerfEval) {
  buildScriptForStirlingPerfEval()
} else {
  buildScriptForCommits()
}
