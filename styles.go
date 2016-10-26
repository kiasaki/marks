package main

const stylesContents = `
/*
background = #fff
color = #111
highlight = #44a4ff
gray-pale = #f8f8f8
gray = #ddd
gray-dark = #999
*/
html, body {
  margin: 0; padding: 0;
  line-height: 1.4;
  font-family: Menlo, Monaco, Lucida Console, Liberation Mono, DejaVu Sans Mono, Bitstream Vera Sans Mono, Courier New, monospace, serif;
  font-size: 16px;
  background: #f8f8f8;
}

* {
  box-sizing: border-box;
}

img {
  max-width: 100%;
}

.btn {
  line-height: 20px;
  padding: 8px 15px;
  border: none;
  background: #44a4ff;
  border: 2px solid #44a4ff;
  color: #fff;
  cursor: pointer;
  display: inline-block;
  font-size: 16px;
}
.btn:hover {
  background: #fff;
  color: #44a4ff;
}
a.btn {
  text-decoration: none;
}
.btn.btn-delete {
  background: #c0392b;
  border-color: #c0392b;
}
.btn.btn-delete:hover {
  background: #fff;
  color: #c0392b;
}
.btn.btn-default {
  background: #999;
  border-color: #999;
}
.btn.btn-default:hover {
  background: #fff;
  color: #999;
}

a {
  color: #44a4ff;
}


/* NAV */
nav {
  background: #fff;
  position: absolute;
  top: 0; bottom: 0; left: 0;
  width: 270px;
  border-right: 1px solid #ddd;
}

/* nav: header */
nav header {
  background: #fff;
  height: 70px;
  padding: 15px;
  line-height: 40px;
  border-bottom: 1px solid #ddd;
}
nav header h1 {
  margin: 0;
  color: #111;
  font-size: 32px;
}
nav header a {
  position: absolute;
  top: 15px; right: 15px;
  line-height: 40px;
  color: #44a4ff;
  font-size: 18px;
}

/* nav: notes */
nav ul {
  padding: 15px 0;
  margin: 0;
  list-style-type: none;
}
nav ul li {
  line-height: 20px;
  font-size: 18px;
}
nav ul li a {
  display: block;
  padding: 15px;
  color: #111;
  text-decoration: none;
  transition: padding 200ms;
  border-left: 5px solid transparent;

  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
nav ul li.active a {
  border-left: 5px solid #44a4ff;
  background: #f8f8f8;
}
nav ul li a:hover {
  background: #f8f8f8;
  border-left: 5px solid #f8f8f8;
}
nav ul li.active a:hover {
  border-left: 5px solid #44a4ff;
}

/* nav: footer */
nav footer {
  position: absolute;
  bottom: 0; left: 0; right: 0;
  background: #fff;
  color: #111;
  padding: 15px;
  border-top: 1px solid #ddd;
}
nav footer a {
  color: #44a4ff;
}


/* PAGE */
section {
  position: absolute;
  top: 0; bottom: 0; left: 270px;
  width: 800px;
  background: #fff;
  overflow: auto;
  border-right: 1px solid #ddd;
}

/* page: header */
section header {
  display: flex;
  height: 70px;
  padding: 15px;
  line-height: 40px;
  border-bottom: 1px solid #ddd;
}
section header input {
  flex: 1;
  height: 40px;
  padding: 8px 0;
  background: transparent;
  border: none;
  line-height: 22px;
  font-size: 22px;
  color: #111;
}
section header input:focus {
  outline: none;
  border-bottom: 1px solid #ddd;
}
section header .buttons {
  flex: 0 0 200px;
  text-align: right;
  height: 40px;
}

/* page: contents */
section .contents textarea {
  width: 100%;
  height: 100%;
  min-height: 800px;
  background: transparent;
  border: none;
}
section .contents .CodeMirror {
  height: calc(100vh - 70px);
}
section .contents .body {
  padding: 15px;
  color: #111;
}
section .contents .body p,
section .contents .body li {
  font-size: 16px;
}
section .contents .body h1,
section .contents .body h2,
section .contents .body h3,
section .contents .body h4,
section .contents .body h5,
section .contents .body h6 {
  margin: 15px 0 0;
}
section .contents .body h1:first-child,
section .contents .body h2:first-child,
section .contents .body h3:first-child,
section .contents .body h4:first-child,
section .contents .body h5:first-child,
section .contents .body h6:first-child {
  margin-top: 0;
}
section .contents .body h1 {
  font-size: 32px;
  padding-bottom: 6px;
  border-bottom: 1px solid #ddd;
}
section .contents .body h2 {
  font-size: 24px;
  padding-bottom: 6px;
  border-bottom: 1px solid #ddd;
}
section .contents .body h3 {
  font-size: 20px;
}
section .contents .body h4 {
  font-size: 16px;
}
section .contents .body blockquote {
  margin-left: 0;
  padding: 0 1em;
  color: #777;
  border-left: 0.25em solid #ddd;
}
section .contents .body pre {
  padding: 15px;
  font-size: 14px;
  overflow: auto;
  background-color: #f8f8f8;
}
`
