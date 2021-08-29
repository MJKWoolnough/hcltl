package main

// File automatically generated with ./embed.sh

const (
	start = "<html lang=\"en\"><head><title>HCL Timeline</title><meta charset=\"UTF-8\" /><link rel=\"shortcut icon\" sizes=\"any\" href=\"data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 256 256'%3E%3Crect x='15' y='15' height='45' width='175' rx='5' fill='%23f00' /%3E%3Crect x='15' y='75' height='45' width='90' rx='5' fill='%230f0' /%3E%3Crect x='160' y='75' height='45' width='85' rx='5' fill='%230f0' /%3E%3Crect x='70' y='135' height='45' width='125' rx='5' fill='%2300f' /%3E%3Crect x='50' y='195' height='45' width='70' rx='5' fill='%23000' /%3E%3Crect x='170' y='195' height='45' width='50' rx='5' fill='%23000' /%3E%3C/svg%3E\" /><script type=\"module\">class t extends HTMLElement{constructor(){super(),new.target===t&&h(this.attachShadow({mode:\"closed\"}),[y({type:\"text/css\"},\":host{display:block;position:relative;overflow:hidden;width:var(--shell-width,100%);height:var(--shell-height,100%)}::slotted(windows-window){--taskmanager-on:none}::slotted(windows-window:last-of-type){--overlay-on:none}\"),b({name:\"desktop\"}),w(b())])}alert(t,e,o){return j(this,t,e,o)}confirm(t,e,o){return U(this,t,e,o)}prompt(t,e,o,i){return F(this,t,e,o,i)}addWindow(t){return this.appendChild(t),!0}realignWindows(){const{offsetWidth:t,offsetHeight:e}=this;Array.from(this.childNodes).filter((t=>t instanceof o)).forEach((o=>{const{offsetLeft:i,offsetTop:s,offsetWidth:n,offsetHeight:r}=o;i+n>t?o.style.setProperty(\"--window-left\",Math.max(t-n,0)+\"px\"):i<0&&o.style.setProperty(\"--window-left\",\"0px\"),s+r>e?o.style.setProperty(\"--window-top\",Math.max(e-r)+\"px\"):s<0&&o.style.setProperty(\"--window-top\",\"0px\")}))}}class e extends HTMLElement{constructor(){super(),this.setAttribute(\"slot\",\"desktop\"),h(this.attachShadow({mode:\"closed\"}),[y({type:\"text/css\"},\":host{position:absolute;top:0;left:0;bottom:0;right:0}\"),b({slot:\"desktop\"})])}}class o extends HTMLElement{constructor(){super();const t=x({part:\"title\"}),e=u({part:\"icon\",src:_}),o=x(),i=()=>this.focus();D.set(this,{titleElement:t,iconElement:e,extraButtons:o}),h(this.attachShadow({mode:\"closed\"}),[y({type:\"text/css\"},':host{position:absolute;display:block;background-color:#fff;color:#000;border:1px solid #000;width:var(--window-width,auto);height:var(--window-height,auto);min-width:100px;min-height:100px;top:var(--window-top,0);left:var(--window-left,0);list-style:none;padding:0;z-index:0}:host([maximised]){left:0;right:0;top:0;bottom:0;width:auto;height:auto}:host([resizable]){border-width:0}:host([resizable])>div:nth-child(2){display:block}:host([minimised]),:host>div:nth-child(2),:host([hide-titlebar])>div:nth-child(3),:host([hide-close])>div:nth-child(3)>div>button:nth-of-type(1),:host([hide-maximise])>div:nth-child(3)>div>button:nth-of-type(2),:host([hide-minimise])>div:nth-child(3)>div>button:nth-of-type(3),:host([window-hide])>div:nth-child(3)>div>button:nth-of-type(3){display:none}:host>div:nth-child(3){white-space:nowrap;height:calc(1em + 6px);background-color:#aaa;overflow:hidden;user-select:none;display:flex;align-items:center}:host>div:nth-child(3)>span{margin-right:calc(3em + 24px)}:host>div:nth-child(3)>img{height:1em;pointer-events:none}:host>div:nth-child(3)>div{position:absolute;right:0;top:0}:host>div:nth-child(3) button{padding:0;border-width:2px;float:right;background-repeat:no-repeat;background-position:center;background-color:#eee;background-size:1em 1em;width:calc(1em + 8px);height:calc(1em + 8px)}:host>div:nth-child(3)>div>button:nth-of-type(1){background-image:url(\\'data:image/svg+xml,%3Csvg viewBox=\"0 0 10 10\" xmlns=\"http://www.w3.org/2000/svg\"%3E%3Cpath d=\"M1,1 L9,9 M9,1 L1,9\" stroke=\"black\" /%3E%3C/svg%3E\\')}:host>div:nth-child(3)>div>button:nth-of-type(2){background-image:url(\\'data:image/svg+xml,%3Csvg viewBox=\"0 0 10 10\" xmlns=\"http://www.w3.org/2000/svg\"%3E%3Cpath d=\"M9,3 h-8 v-1 h8 v-1 h-8 v8 h8 v-8\" stroke=\"black\" fill=\"none\" /%3E%3C/svg%3E\\')}:host([maximised])>div:nth-child(3)>div>button:nth-of-type(2){background-image:url(\\'data:image/svg+xml,%3Csvg viewBox=\"0 0 15 13\" xmlns=\"http://www.w3.org/2000/svg\"%3E%3Cpath d=\"M1,5 h8 v-1 h-8 v8 h8 v-8 m-3,0 v-3 h8 v8 h-5 m5,-7 h-8\" stroke=\"%23000\" fill=\"none\" /%3E%3C/svg%3E\\')}:host>div:nth-child(3)>div>button:nth-of-type(3){display:var(--taskmanager-on,block);background-image:url(\\'data:image/svg+xml,%3Csvg viewBox=\"0 0 10 10\" xmlns=\"http://www.w3.org/2000/svg\"%3E%3Cline x1=\"1\" y1=\"9\" x2=\"9\" y2=\"9\" stroke=\"black\" /%3E%3C/svg%3E\\')}:host>div:nth-child(4){user-select:contain}:host([resizable])>div:nth-child(4){overflow:auto;position:absolute;bottom:0;left:0;right:0;top:calc(1em + 6px)}:host>div:nth-child(5){pointer-events:none;display:var(--overlay-on,block);position:absolute;background-color:rgba(0,0,0,0.1);top:0;left:0;bottom:0;right:0}:host([resizeable])>div:nth-child(5){top:var(calc(--window-resize), -2px);left:var(calc(--window-resize), -2px);bottom:var(calc(--window-resize), -2px);right:var(calc(--window-resize), -2px)}:host>div:nth-child(2)>div{position:absolute;border-color:currentColor;border-style:solid;border-width:0;z-index:-1}:host>div:nth-child(2)>div:nth-child(1){top:-2px;left:-2px;width:10px;height:10px;cursor:nwse-resize;border-left-width:3px;border-top-width:3px}:host>div:nth-child(2)>div:nth-child(2){top:-2px;left:8px;right:8px;border-top-width:3px;cursor:ns-resize}:host>div:nth-child(2)>div:nth-child(3){top:-2px;right:-2px;width:10px;height:10px;border-top-width:3px;border-right-width:3px;cursor:nesw-resize}:host>div:nth-child(2)>div:nth-child(4){top:8px;right:-2px;bottom:8px;border-right-width:3px;cursor:ew-resize}:host>div:nth-child(2)>div:nth-child(5){bottom:-2px;right:-2px;width:10px;height:10px;border-right-width:3px;border-bottom-width:3px;cursor:nwse-resize}:host>div:nth-child(2)>div:nth-child(6){bottom:-2px;left:8px;right:8px;border-bottom-width:3px;cursor:ns-resize}:host>div:nth-child(2)>div:nth-child(7){bottom:-2px;left:-2px;width:10px;height:10px;border-left-width:3px;border-bottom-width:3px;cursor:nesw-resize}:host>div:nth-child(2)>div:nth-child(8){top:8px;left:-2px;bottom:8px;border-left-width:3px;cursor:ew-resize}'),w(Array.from({length:8},((t,e)=>w({onmousedown:O.bind(this,e)})))),w({part:\"titlebar\",onmousedown:V.bind(this),ondblclick:t=>{t.target instanceof HTMLButtonElement||this.hasAttribute(\"hide-maximise\")||this.toggleAttribute(\"maximised\")}},[e,t,w({part:\"controls\"},[f({part:\"close\",onclick:()=>this.close()}),f({part:\"maximise\",onclick:()=>this.toggleAttribute(\"maximised\")}),f({part:\"minimise\",onclick:()=>this.toggleAttribute(\"minimised\")}),o])]),w(b()),w({onclick:i})]),this.addEventListener(\"mousedown\",i,{capture:!0})}connectedCallback(){i===this&&(i=null)}disconnectedCallback(){if(i===this)return;const t=Y.get(this),e=X.get(this);t&&(X.delete(t),Y.delete(this)),e&&(e.remove(),X.delete(this)),this.dispatchEvent(new CustomEvent(\"remove\",{cancelable:!1}))}attributeChangedCallback(t,e,o){const i=D.get(this);switch(t){case\"window-title\":i.titleElement.innerText=o;break;case\"window-icon\":i.iconElement.setAttribute(\"src\",o)}}static get observedAttributes(){return[\"window-title\",\"window-icon\"]}alert(t,e,o){return j(this,t,e,o)}confirm(t,e,o){return U(this,t,e,o)}prompt(t,e,o,i){return F(this,t,e,o,i)}addWindow(t){return!!this.parentNode&&(X.has(this)?(X.get(this).addWindow(t),!0):(X.set(this,t),Y.set(t,this),this.parentNode.appendChild(t),!0))}addControlButton(t,e,o){const i=D.get(this).extraButtons.appendChild(f({style:{\"background-image\":`url(${t})`},onclick:()=>e.call(this),title:o}));return()=>i.remove()}focus(){const t=X.get(this);t?t.focus():(this.parentNode&&this.nextElementSibling&&(i=this,this.parentNode.appendChild(this)),super.focus())}close(){if(X.has(this))X.get(this).focus();else if(this.dispatchEvent(new CustomEvent(\"close\",{cancelable:!0})))return this.remove(),!0;return!1}}customElements.define(\"windows-shell\",t),customElements.define(\"windows-desktop\",e),customElements.define(\"windows-window\",o);let i=null,s=!1,n=20;const r=(t,e)=>{if(\"string\"==typeof e)t.appendChild(document.createTextNode(e));else if(Array.isArray(e))for(const o of e)r(t,o);else if(e instanceof Node)t.appendChild(e);else if(e instanceof NodeList)for(const o of e)t.appendChild(o)},l=t=>(e,o,i)=>{const s=\"string\"==typeof e?document.createElementNS(t,e):e instanceof Node?e:document.createDocumentFragment();if(!(\"string\"==typeof o||o instanceof Array||o instanceof NodeList||o instanceof Node)&&(\"object\"!=typeof i||i instanceof Array||i instanceof Node||i instanceof NodeList)||([o,i]=[i,o]),\"object\"==typeof o&&s instanceof Element)for(const[t,e]of Object.entries(o))if(e instanceof Function){const o={};let i=t;t:for(;;){switch(i.charAt(0)){case\"1\":o.once=!0;break;case\"C\":o.capture=!0;break;case\"P\":o.passive=!0;break;default:break t}i=i.slice(1)}i.startsWith(\"on\")&&s.addEventListener(i.substr(2),e,o)}else if(\"class\"===t&&(e instanceof Array||e instanceof DOMTokenList)&&e.length>0)s.classList.add(...e);else if(\"style\"===t&&\"object\"==typeof e&&(s instanceof HTMLElement||s instanceof SVGElement))for(const t in e)void 0===e[t]?s.style.removeProperty(t):s.style.setProperty(t,e[t]);else\"string\"==typeof e||\"number\"==typeof e?s.setAttribute(t,e):\"boolean\"==typeof e?s.toggleAttribute(t,e):void 0===e&&s.hasAttribute(t)&&s.removeAttribute(t);return\"string\"==typeof i?s.textContent=i:i&&(i instanceof Array||i instanceof Node||i instanceof NodeList)&&r(s,i),s},h=l(\"http://www.w3.org/1999/xhtml\"),d=l(\"http://www.w3.org/2000/svg\"),a=t=>{for(;null!==t.lastChild;)t.removeChild(t.lastChild);return t},c=(t,e=!0)=>(window.setTimeout((()=>{t.focus(),(t instanceof HTMLInputElement||t instanceof HTMLTextAreaElement)&&e&&t.select()}),0),t),[p,f,w,m,u,g,v,x,b,y,k,E,C,$,L,P]=\"br button div h1 img input label span slot style table tbody td th thead tr\".split(\" \").map((t=>h.bind(null,t))),[M,T,A,z,W,H,N,S,B,I]=\"circle defs g line path pattern rect svg text use\".split(\" \").map((t=>d.bind(null,t))),D=new WeakMap,O=function(e,o){const i=this.parentNode;if(s||!(i instanceof t)||0!==o.button)return;s=!0,i.style.setProperty(\"user-select\",\"none\");const n=this.offsetLeft,r=this.offsetTop,l=this.offsetWidth,h=this.offsetHeight,d=o.clientX,a=o.clientY,c=t=>{const o=t.clientX-d,i=t.clientY-a;switch(e){case 0:case 1:case 2:{const t=h-i;t>(parseInt(this.style.getPropertyValue(\"min-height\")||\"\")||100)&&(this.style.setProperty(\"--window-top\",`${r+i}px`),this.style.setProperty(\"--window-height\",`${t}px`))}break;case 4:case 5:case 6:{const t=h+i;t>(parseInt(this.style.getPropertyValue(\"min-height\")||\"\")||100)&&this.style.setProperty(\"--window-height\",`${t}px`)}}switch(e){case 2:case 3:case 4:{const t=l+o;t>(parseInt(this.style.getPropertyValue(\"min-width\")||\"\")||100)&&this.style.setProperty(\"--window-width\",`${t}px`)}break;case 0:case 7:case 6:{const t=l-o;t>(parseInt(this.style.getPropertyValue(\"min-width\")||\"\")||100)&&(this.style.setProperty(\"--window-left\",`${n+o}px`),this.style.setProperty(\"--window-width\",`${t}px`))}}},p=t=>{0===t.button&&(i.removeEventListener(\"mousemove\",c),i.removeEventListener(\"mouseup\",p),i.style.removeProperty(\"user-select\"),s=!1,this.dispatchEvent(new CustomEvent(\"resized\")))};i.addEventListener(\"mousemove\",c),i.addEventListener(\"mouseup\",p)},V=function(e){const i=this.parentNode;if(s||!(i instanceof t)||0!==e.button)return;s=!0,this.style.setProperty(\"user-select\",\"none\");const n=e.clientX-this.offsetLeft,r=e.clientY-this.offsetTop,l=t=>{const e=t.clientX-n,s=t.clientY-r,[l,h]=((t,e,i,s)=>{const n=parseInt(t.getAttribute(\"snap\")||\"0\"),{offsetLeft:r,offsetTop:l,offsetWidth:h,offsetHeight:d}=e,a=r+h,c=l+d,p=i+h,f=s+d,w=[0,0];return n>0&&(r>=0&&i<0&&i>=-n?w[0]=-i:a<=t.offsetWidth&&p>t.offsetWidth&&p<=t.offsetWidth+n&&(w[0]=t.offsetWidth-p),l>=0&&s<0&&s>=-n?w[1]=-s:c<=t.offsetHeight&&f>t.offsetHeight&&f<=t.offsetHeight+n&&(w[1]=t.offsetHeight-f),Array.from(t.childNodes).filter((t=>t instanceof o&&t!==e)).forEach((t=>{const e=t.offsetLeft,o=t.offsetTop,h=e+t.offsetWidth,d=o+t.offsetHeight;s<=d&&f>=o&&(a<=e&&p>=e&&p<=e+n?w[0]=e-p:r>=h&&i<=h&&i>=h-n&&(w[0]=h-i)),i<=h&&p>=e&&(c<=o&&f>=o&&f<=o+n?w[1]=o-f:l>=d&&s<=d&&s>=d-n&&(w[1]=d-s))}))),w})(i,this,e,s);this.style.setProperty(\"--window-left\",e+l+\"px\"),this.style.setProperty(\"--window-top\",s+h+\"px\")},h=t=>{0===t.button&&(i.removeEventListener(\"mousemove\",l),i.removeEventListener(\"mouseup\",h),i.style.removeProperty(\"user-select\"),s=!1,this.dispatchEvent(new CustomEvent(\"moved\")))};i.addEventListener(\"mousemove\",l),i.addEventListener(\"mouseup\",h)},X=new Map,Y=new Map,j=(t,e,o,i)=>new Promise(((s,n)=>{const r=R({\"window-hide\":!0,\"window-icon\":i||_,\"window-title\":e,\"hide-maximise\":\"true\",onremove:()=>s(!1)},[w(o),w({style:\"text-align: center\"},c(f({onclick:()=>{s(!0),r.remove()}},\"Ok\")))]);t.addWindow(r)||n(new Error(\"invalid target\"))})),U=(t,e,o,i)=>new Promise(((s,n)=>{const r=R({\"window-hide\":!0,\"window-icon\":i||_,\"window-title\":e,\"hide-maximise\":\"true\",onremove:()=>s(!1)},[w(o),w({style:\"text-align: center\"},[c(f({onclick:()=>{s(!0),r.remove()}},\"Ok\")),f({onclick:()=>r.remove()},\"Cancel\")])]);t.addWindow(r)||n(new Error(\"invalid target\"))})),F=(t,e,o,i,s)=>new Promise(((n,r)=>{const l=f({onclick:()=>{n(h.value),d.remove()}},\"Ok\"),h=c(g({value:i||\"\",onkeydown:t=>{switch(t.key){case\"Enter\":l.click();break;case\"Escape\":d.remove()}}})),d=R({\"window-hide\":!0,\"window-icon\":s||_,\"window-title\":e,\"hide-maximise\":\"true\",onremove:()=>n(null)},[w(o),h,w({style:\"text-align: center\"},l)]);t.addWindow(d)||r(new Error(\"invalid target\"))})),R=(t,e)=>h(new o,t,e),_='data:image/svg+xml,%3Csvg viewBox=\"0 0 14 18\" xmlns=\"http://www.w3.org/2000/svg\"%3E%3Cpath d=\"M9,1 h-8 v16 h12 v-12 Z v4 h4\" stroke=\"black\" fill=\"none\" /%3E%3Cpath d=\"M3,8 h8 m-8,3 h8 m-8,3 h8\" stroke=\"gray\" /%3E%3C/svg%3E',G=["
	mid   = "],Z=["
	mid2  = "],q=["
	end   = "],J=Array.from({length:Z.length},(()=>!0)),K=[\"#29f\",\"#f43\",\"#fe3\",\"#4a5\",\"#92b\",\"#0bd\",\"#f52\",\"#cd3\",\"#35b\",\"#e16\",\"#098\"],Q=[[30,\"#0f0\"],[80,\"#ff0\"],[240,\"#fa0\"],[1/0,\"#f00\"]],tt=w({id:\"mouseLine\"}),et=w({id:\"mouseLine2\"}),ot=w({style:{\"background-color\":\"#fff\",border:\"1px solid #000\",position:\"absolute\",top:0,\"z-index\":10}}),it=w(),st=h(new t,void 0,void 0),nt=t=>{const e=t+\"\";return 1==e.length?\"0\"+e:e},rt=t=>{const e=new Date(1e3*t);return`${e.getFullYear()}-${nt(e.getMonth()+1)}-${e.getDate()} ${nt(e.getHours())}:${nt(e.getMinutes())}:${nt(e.getSeconds())}`},lt=t=>`${Math.floor(t/3600)}:${nt(Math.floor(t%3600/60))}:${nt(t%60)}`,ht=S({viewBox:\"0 0 20 20\",onclick:()=>{st.addWindow(at),at.focus()}},[T(W({id:\"spoke\",d:\"M1,7 v2 a1,1 0,0,1 -2,0 v-2 z\",fill:\"#aaa\"})),A({transform:\"translate(10, 10)\"},[M({r:5.5,fill:\"none\",stroke:\"#aaa\",\"stroke-width\":4.5}),Array.from({length:12},((t,e)=>30*e)).map((t=>I({href:\"#spoke\",transform:`rotate(${t})`})))])]),dt=$({rowspan:2},w(ht)),at=R({\"window-icon\":\"data:image/svg+xml,\"+encodeURIComponent('<svg xmlns=\"http://www.w3.org/2000/svg\"'+ht.outerHTML.slice(4).replaceAll(\"#aaa\",\"#000\")),\"window-title\":\"Settings\",tabindex:-1,onkeydown:t=>{console.log(t.key),\"Escape\"===t.key&&at.close()}},[m(\"Settings\"),v({for:\"scale\"},\"Timeline Scale (pixels per minute): \"),g({id:\"scale\",type:\"number\",min:1,value:n,onchange:function(){n=parseInt(this.value)}}),p(),p(),w({style:{\"text-decoration\":\"underline\"}},\"Toggle Users:\"),p(),Z.map(((t,e)=>[v({for:`user_${e}`},`${t}: `),g({type:\"checkbox\",id:`user_${e}`,checked:!0,onclick:function(){J[e]=this.checked}}),p()])),f({onclick:()=>ct(G)},\"Rebuild\")]),ct=t=>{const e=new Map,o=E(),i=[],s=t=>{const e=t.offsetX+(t.target instanceof HTMLDivElement?t.target.offsetLeft:0);tt.style.setProperty(\"left\",dt.offsetWidth+e+\"px\"),et.style.setProperty(\"left\",e+\"px\"),ot.style.setProperty(\"left\",dt.offsetWidth+e+\"\"),ot.innerText=rt(l+60*(e+2)/n)},r=[];let l=1/0,d=-1/0,c=0;for(const n of t){const[t,r,h,a]=n;if(!J[t])continue;let c,p,f,m=!1;if(e.has(t))[,,c,p,f]=e.get(t);else{const i=w(Z[t]),n=C({style:{color:K[e.size%K.length]},onmousemove:s});e.set(t,[i,n,c=[],p=[-1,0],f=[]]),o.appendChild(P([$(i),n]))}r>p[0]?(p[0]>-1&&f.push([p[0],r]),p[1]+=h-r,p[0]=h):h>p[0]&&(p[1]+=h-p[0],p[0]=h);t:for(const t of c){for(const[,e,o]of t)if(r<o&&h>e)continue t;m=!0,t.push(n);break}if(m||(c.push([n]),m=!1),a>0){t:for(const t of i){for(const[,e,,o]of t)if(a<e&&r>o)continue t;m=!0,t.push(n);break}m||i.push([n]),a<l&&(l=a)}r<l&&(l=r),h>d&&(d=h)}l=60*Math.floor(l/60),d=60*Math.ceil(d/60);for(const[,[t,o,i,[,s],r]]of e){let e=0,h=0,d=0;for(const t of i){for(const[,i,s,,r]of t)o.appendChild(w({title:`${rt(i)} ⟶  ${rt(s)}\\nCall Time: ${lt(s-i)}${r?`\\n${q[r]}`:\"\"}`,style:{width:n*(s-i)/60+1+\"px\",left:n*(i-l)/60-2+\"px\",\"--row\":e}})),h++,d+=s-i;e++}for(const[t,e]of r)o.appendChild(w({class:\"wait\",title:`Down Time: ${lt(e-t)}`,style:{left:n*(t-l)/60-2+\"px\",width:n*(e-t)/60+1+\"px\"}}));t.setAttribute(\"title\",`Total Calls: ${h}\\nTotal Call Time: ${lt(d)}\\nNon-concurrent Call Time: ${lt(s)}`),c=Math.max(i.length,c)}tt.style.setProperty(\"--h\",c*e.size+i.length+\"\"),document.body.style.setProperty(\"--rows\",c+\"\");for(let t=3600*Math.ceil(l/3600);t<=d;t+=3600){const e=new Date(1e3*t);r.push(B({x:n*(t-l)/60,y:15},`${nt(e.getHours())}:00`))}ot.style.setProperty(\"left\",\"-1000px\"),h(a(it),k([L({style:{\"--rows\":i.length}},[P([dt,C({style:{width:n*(d-l)/60+\"px\"}},[S({width:n*(d-l)/60,height:20,viewBox:`0 0 ${n*(d-l)/60} 20`},[T(H({id:\"ticks\",patternUnits:\"userSpaceOnUse\",width:60*n,height:20,x:-l%3600/60*n-2},[z({y2:20,stroke:\"#000\"}),z({x1:15*n,x2:15*n,y2:10,stroke:\"#000\"}),z({x1:30*n,x2:30*n,y2:10,stroke:\"#000\"}),z({x1:45*n,x2:45*n,y2:10,stroke:\"#000\"})])),N({width:\"100%\",height:\"100%\",fill:\"url(#ticks)\"}),r]),ot])]),P(C({onmousemove:s},[et,i.map(((t,e)=>t.map((([t,o,,i,s])=>w({title:`${Z[t]}\\n${rt(i)} ⟶  ${rt(o)}\\nWait Time: ${lt(o-i)}${s?`\\n${q[s]}`:\"\"}`,style:{width:n*(o-i)/60+1+\"px\",left:n*(i-l)/60-2+\"px\",\"--row\":e,color:Q.find((([t])=>t>o-i))[1]}})))))]))]),o]))};G.sort((([,t],[,e])=>t-e)),(\"complete\"==document.readyState?Promise.resolve():new Promise((t=>globalThis.addEventListener(\"load\",t,{once:!0})))).then((()=>{h(a(document.body),[it,tt,st]),ct(G)}));</script><style type=\"text/css\">html,body{margin:0;background-color:#fff;user-select:none}text{text-anchor:middle}table{table-layout:fixed;border-collapse:collapse;border-spacing:0;margin:0;white-space:nowrap;padding:0}td,th{padding:0;margin:0}thead th{text-align:center;background-color:#fff}thead{position:sticky;top:0;z-index:30}thead{background-color:#fff}thead th svg{width:1em;cursor:pointer}thead td:first-child{background-color:#eee}thead td:first-child,tbody td{position:relative;height:calc(var(--rows) * 2em)}thead td:first-child div,tbody td div{position:absolute;top:calc(var(--row) * 2em + 3px);height:calc(2em - 6px);background-color:currentColor;z-index:10}tbody td div{border:1px solid #000;box-sizing:border-box}tbody tr th{background-color:#ff0}tbody tr:nth-child(2n) th{background-color:#ee0}tbody tr:nth-child(2n){background-color:#ddd}th{position:sticky;left:0;z-index:20}.wait{top:0;height:calc(var(--rows) * 2em);border-width:0;background-color:transparent}#mouseLine,#mouseLine2{border-left:1px solid red;position:absolute;top:1px;z-index:0;left:-10px;height:calc(2em * var(--h) + 21px)}#mouseLine2{height:calc(2em * var(--rows))}windows-shell{height:0}windows-window{z-index:40;position:fixed;outline:0}</style></head><body></body></html>"
)
