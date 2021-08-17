import type {Data} from './data.js';
import {createHTML, clearElement} from './lib/dom.js';
import {br, button, div, h1, input, label, table, tbody, td, th, thead, tr} from './lib/html.js';
import {circle, defs, g, line, path, pattern, rect, svg, text, use} from './lib/svg.js';
import {windows, shell} from './lib/windows.js';
import {data, users} from './data.js';

declare const pageLoad: Promise<void>;

let minuteWidth = 20;

const userFilter = Array.from({"length": users.length}, () => true),
      colours = [
	"#29f",
	"#f43",
	"#fe3",
	"#4a5",
	"#92b",
	"#0bd",
	"#f52",
	"#cd3",
	"#35b",
	"#e16",
	"#098",
      ],
      ml = div({"id": "mouseLine"}),
      mt = div({"style": {"background-color": "#fff", "border": "1px solid #000", "position": "absolute", "top": 0, "z-index": 10}}, "TIME"),
      timeline = div(),
      s = shell(),
      pad = (n: number) => {
	const t = n + "";
	return t.length == 1 ? "0" + t : t;
      },
      formatTime = (time: number) => {
	const d = new Date(time * 1000);
	return `${d.getFullYear()}-${pad(d.getMonth()+1)}-${d.getDate()} ${pad(d.getHours())}:${pad(d.getMinutes())}:${pad(d.getSeconds())}`;
      },
      settings = td(div(svg({"viewBox": "0 0 20 20", "onclick": () => s.addWindow(settingsWindow)}, [
	defs(path({"id": "spoke", "d": "M1,7 v2 a1,1 0,0,1 -2,0 v-2 z", "fill": "#aaa"})),
	g({"transform": "translate(10, 10)"}, [
		circle({"r": 5.5, "fill": "none", "stroke": "#aaa", "stroke-width": 4.5}),
		Array.from({"length": 12}, (_, n) => n * 30).map(r => use({"href": "#spoke", "transform": `rotate(${r})`})),
	])
      ]))),
      settingsWindow = windows({"window-icon": "data:image/svg+xml," + encodeURIComponent("<svg xmlns=\"http://www.w3.org/2000/svg\"" + settings.outerHTML.slice(8, -5).replaceAll("#aaa", "#000")), "window-title": "Settings"}, [
	      h1("Settings"),
	      label({"for": "scale"}, "Timeline Scale (pixels per minute): "),
	      input({"id": "scale", "type": "number", "min": 1, "value": minuteWidth, "onchange": function(this: HTMLInputElement) {minuteWidth = parseInt(this.value);}}),
	      br(),
	      br(),
	      div({"style": {"text-decoration": "underline"}}, "Toggle Users:"),
	      br(),
	      users.map((user, n) => [
		label({"for": `user_${n}`}, `${user}: `),
		input({"type": "checkbox", "id": `user_${n}`, "checked": true, "onclick": function(this: HTMLInputElement) {userFilter[n] = this.checked}}),
		br(),
	      ]),
	      button({"onclick": () => buildTimeline(data)}, "Rebuild")
      ]),
      buildTimeline = (data: Data) => {
	const rows = new Map<number, [HTMLDivElement, HTMLDivElement, Data[]]>(),
	      tb = tbody(),
	      mm = (e: MouseEvent) => {
		const offset = e.offsetX + (e.target instanceof HTMLDivElement ? e.target.offsetLeft : 0);
		ml.style.setProperty("left", settings.offsetWidth + offset + "px");
		mt.style.setProperty("left", (offset - 1) + "");
		mt.innerText = formatTime(earliest + 60 * (offset+1) / minuteWidth);
	      },
	      loggedRows: Data[] = [];
	let earliest = Infinity,
	    latest = -Infinity,
	    maxRows = 0;
	for (const row of data) {
		const [user, start, stop, logged] = row;
		if (!userFilter[user]) {
			continue;
		}
		let d: Data[],
		    set = false;
		if (!rows.has(user)) {
			const h = div(users[user]),
			      t = td({"style": {"color": colours[rows.size % colours.length]}, "onmousemove": mm});
			rows.set(user, [h, t, d = []]);
			tb.appendChild(tr([th(h), t]));
		} else {
			[, , d] = rows.get(user)!;
		}
		DLoop:
		for (const r of d) {
			for (const [, cstart, cstop] of r) {
				if (start < cstop && stop > cstart) {
					continue DLoop
				}
			}
			set = true;
			r.push(row);
			break;
		}
		if (!set) {
			d.push([row])
			set = false;
		}
		LLoop:
		for (const r of loggedRows) {
			for (const [, cstop,, cstart] of r) {
				if (logged < cstop && start > cstart) {
					continue LLoop
				}
			}
			set = true;
			r.push(row);
			break;
		}
		if (!set) {
			loggedRows.push([row])
		}
		if (logged < earliest) {
			earliest = start;
		}
		if (stop > latest) {
			latest = stop;
		}
	}
	earliest = Math.floor(earliest / 60) * 60;
	latest = Math.ceil(latest / 60) * 60;
	for (const [, [t, cell, d]] of rows) {
		let rnum = 0,
		    calls = 0,
		    secs = 0;
		for (const row of d) {
			for (const [, start, stop] of row) {
				cell.appendChild(div({"title": formatTime(start) + " ⟶ " + formatTime(stop), "style": {"width": (minuteWidth * (stop - start) / 60 + 1) + "px", "left": (minuteWidth * (start - earliest) / 60 - 2) + "px", "--row": rnum}}));
				calls++;
				secs += stop - start;
			}
			rnum++;
		}
		t.setAttribute("title", `Total Calls: ${calls}\nTotal Call Time: ${Math.floor(secs / 3600)}:${pad(Math.floor((secs % 3600) / 60))}:${pad(secs % 60)}`);
		maxRows = Math.max(d.length, maxRows);
	}
	ml.style.setProperty("--h",  (maxRows * rows.size + loggedRows.length) + "");
	document.body.style.setProperty("--rows", maxRows+"");
	tb.appendChild(tr([
		th({"style": {"--rows": loggedRows.length}}, div()),
		td({"onmousemove": mm}, loggedRows.map((row, n) => row.map(([user, start,, logged]) => div({"title": `${users[user]}\n${formatTime(logged)} ⟶  ${formatTime(start)}`, "style": {"width": (minuteWidth * (start - logged) / 60 + 1) + "px", "left": (minuteWidth * (logged - earliest) / 60 - 2) + "px" , "--row": n}}))))
	]));
	const hours: SVGTextElement[] = [];
	for (let hour = Math.ceil(earliest / 3600) * 3600; hour <= latest; hour += 3600) {
		const d = new Date(hour * 1000);
		hours.push(text({"x": (minuteWidth * (hour - earliest) / 60), "y": 15}, `${pad(d.getHours())}:00`));
	}
	mt.style.setProperty("left", "-1000px");
	createHTML(clearElement(timeline), table([
		thead(tr([settings, th({"style": {"width": (minuteWidth * (latest - earliest) / 60) + "px"}}, [
			svg({"width": minuteWidth * (latest - earliest) / 60, "height": 20, "viewBox": `0 0 ${minuteWidth * (latest - earliest) / 60} 20`}, [
				defs(pattern({"id": "ticks", "patternUnits": "userSpaceOnUse", "width": 60 * minuteWidth, "height": 20, "x": -(earliest % 3600) / 60 * minuteWidth - 2}, [
					line({"y2": 20, "stroke": "#000"}),
					line({"x1": minuteWidth * 15, "x2": minuteWidth * 15, "y2": 10, "stroke": "#000"}),
					line({"x1": minuteWidth * 30, "x2": minuteWidth * 30, "y2": 10, "stroke": "#000"}),
					line({"x1": minuteWidth * 45, "x2": minuteWidth * 45, "y2": 10, "stroke": "#000"})
				])),
				rect({"width": "100%", "height": "100%", "fill": "url(#ticks)"}),
				hours,
			]),
			mt
		])])),
		tb,
	]));
      };

data.sort(([,a], [, b]) => a - b);

pageLoad.then(() => {
	createHTML(clearElement(document.body), [
		timeline,
		ml,
		s,
	]);
	buildTimeline(data);
});
