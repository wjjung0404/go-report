# Vue 설치 및 실행

1. Command (cmd)를 열어 cd 명령어를 통해 프로젝트를 만들 폴더로 이동

> cd <프로젝트 생성주소>

2. npm 또는 yarn으로 vue 설치
npm : npm install --global @vue/cli
yarn : yarn global add @vue/cli

>   node.js 설치가 되어있지 않으면 npm 또는 yarn 사용 불가.
node.js 설치를 먼저 진행해주세요.

3. 설치후 다음 명령어 실행
   vue create <프로젝트 이름>

> 1. Manually select features 옵션을 선택하면 설치가능한 vue 라이브러리를 선택할 수 있습니다.
>    필요에 따라 선택해서 설치해주세요.
> 2. 해당예시는 Babel과 Linter / Formatter 만 설치합니다.
>    Linter / Formatter 설치시 버전 선택후 config 설정시
> 3. ESlint wtih error prevention only 선택합니다.
>    자동 lint를 어떤 방식으로 하고 싶은지에 대한 메세지가 나옵니다.
> 4. Lint on save로 선택
>    구성파일 관리에 대한 선택이 나옵니다.
> 5. In package.json옵션은 모든 구성 설정을 앱의 package.json에 저장합니다.
>    해당 예에서는 In dedicated config files 옵션을 선택합니다.
> 6. 마지막으로 설정한 내용을 다음에 재사용할지 여부입니다.
>    필요에 따라 설정해주세요.
> * 모든 옵션에 대한 설정은 해당 페이지에서 참고하세요
> <a href="https://cli.vuejs.org/">find more information on the CLI</a>

<h2>파일에 대한 설명</h2>
<b>.eslintrc.js</b> : eslint의 config 파일. linting 규칙관리
<b>.babel.config.js</b> : Babel config 파일. 최신 JS 기능을 상용 브라우저와 호환 가능하도록 구버전으로 번역 (호환시킴)
<b>.browserslistrc</b> : Browserslist config 파일. 도구(툴)을 어떤 브라우저에 최적화할지 정의할 수 있습니다..
<b>public 디렉토리</b> : Webpack의 처리를 받지않고 퍼블리싱되는 정적 에셋(Static assets) 포함.(index.html의 경우 일부 Webpack처리 받음)
<b>public\favico.ico</b> : 앱의 대표 favicon 파일. Vue 로고가 기본값으로 지정되어 있습니다.
<b>public\index.html</b> : 앱의 템플릿 파일. Vue앱은 이 HTML페이지로부터 실행됨. lodash 템플릿 구문은 사용해 보간법으로 값을 지정할 수 있습니다.
<b>src</b> : Vue 애플리케이션 핵심 내용을 포함하게 될 디렉토리
<b>src\main.js</b> : 애플리케이션 진입점. main.js는 Vue앱을 초기화하고 index.html 파일에서 어떤 HTML엘리멘트를 붙여야 하는지(#app 엘리멘트) 나타냅니다.
이 파일에 전역 컴포넌트나 부가적인 Vue 라이브러리를 등록하는 경우가 많습니다.
<b>src\App.vue</b> : Vue 애플리케이션의 최상위 컴포넌트.
<b>src\components</b> : Vue 컴포넌트를 저장할 디렉토리.
<b>src\assets</b> : CSS, 이미지 등의 정적자산(static assets) 저장하는 디렉토리.
src 디렉토리에 포함되어 있어 Webpack 처리를 받습니다.(SCSS/SCSS 또는 Stylus 전처리기 도구를 사용할 수 있다.)

<h2>.vue 파일(단일 파일 컴포넌트)</h2>
Vue앱을 구축할 때는 컴포넌트가 중심적인 역할을 수행합니다.
여러 개의 개별 컴포넌트로 분리해 별도로 생성/구축후 필요에 따라 컴포넌트간 데이터를 주고 받을 수 있으며,
작은 블록들은 우리가 코드에 대해 추론하고 테스트하기 쉽도록 도와줍니다.

Vue는 단일 파일 컴포넌트를 사용하면 템플릿과 대응하는 스크립트, CSS를 하나의 .vue 파일로 묶어서 작성하도록 하며, 이 파일들을 Webpack과 같은 JS 빌드 도구를 통해 처리합니다. 
(빌드 타임 도구를 이용가능)
<b>Bable, TypeScript, SCSS 같은 도구를 사용하여 정교한 컴포넌트 생성 가능</b>

CLI로 생성한 프로젝트는 Webpack과 함게 .vue 파일을 사용할 수 있도록 구성합니다.
src 폴더에 App.vue가 생성된것을 확인할 수 있습니다.

<h2>App.vue</h2>
<b>App.vue</b>파일에는 크게 세부분으로 이루어져 있습니다.
1. <b><i><template></i> 파트</b> : 컴포넌트 템플릿 정의
2. <b><i><script></i> 파트</b> : 스크립트를 작성
3. <b><i><style></i> 파트</b> : 스타일시트 파트
<script>는 컴포넌트 화면에 표시되지 않는 모든 로직을 포함하고 있습니다.
<script>태그 안에 반드시 기본으로 export(export default 구문 참고)되는 JS 오브젝트가 있어야 합니다.
> 오브젝트에선 로컬 컴포넌트 등록, 컴포넌트 인풋(props) 정의, 로컬 상태 관리, > 메서드 정의 등의 작업이 이루어집니다.
> 빌드 단계에서 이 오브젝트가 처리되고, 템플릿과 함께 render()함수를 통해 vue > 컴포넌트로 변환됩니다.

<b>default export 오브젝트</b> 오브젝트는 컴포넌트 이름을 app으로 설정하고
components 속성에 HelloWorld.vue 컴포넌트를 등록했습니다.
이렇게 컴포넌트를 등록하면 로컬 컴포넌트가 됩니다.
> 로컬로 등록된 컴포넌트는 하위 컴포넌트로 사용이 불가합니다. 그러므로 각 컴포넌트 파일에 필요한 컴포넌트를 import 하고 등록해야 합니다.
(번들 분할/트리 쉐이킹에 유용한 기능.)

<b><style></b>에는 컴포넌트에 사용될 CSS를 포함합니다.
<style scoped>같이 scoped속성 추가시 Vue는 그 안의 내용을 단일 파일 컴포넌트(SFC) 내부 범위에서만 적용합니다. (CSS-in-JS 방식과 비슷하게 동작, 일반 CSS 구문 작성할 수 있음.)

<h2>로컬에서 앱 실행</h2>
Vue CLI는 개발 서버를 내장하고 있습니다.
터미널(CMD)를 통해 해당 폴더로 이동하고 npm run serve 또는 yarn serve 를 실행해보세요.
> 서버 종료는 Ctrl + C 를 통해 종료할 수 있습니다.

<h2>실습. App.vue 변경</h2>
App.vue 파일을 열어 템플릿 섹션에서 <img> 요소를 지워봅시다.
> <img alt="Vue logo" src="./assets/logo.png">
서버를 실행중이라면 로고가 즉시 사라진것을 확인할 수 있습니다.

템플릿 섹션의 HelloWorld 컴포넌트를 삭제해봅시다.
> 1. <HelloWorld msg="Welcome to Your Vue.js App"/>
이제 App.vue파일을 저장하면 HelloWorld 컴포넌트를 등록했으나 사용하지 않았다는 에러를 발생시킵니다. (등록한 컴포넌트는 반드시 사용해주기)

해당 라인도 삭제해봅시다.
> import HelloWorld from './components/HelloWorld.vue'
> components: {
>  	HelloWorld
>}
에러가 발생시키는 요소를 지웠습니다. 
<template> 파트에 표시할 내용이 없기 떄문에 빈 페이지만 보입니다.

<b>요소 추가</b>
<div id="app"> 안에 h1 요소를 추가해보겠습니다.
><template>
>  <div id="app">
>    <h1>To-Do List</h1>
>  </div>
></template>
헤더 텍스트를 To-Do List로 작성해줍니다.

