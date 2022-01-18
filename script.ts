import type {Data} from './data.js';
import {amendNode, clearNode} from './lib/dom.js';
import {br, button, div, h1, input, label, table, tbody, td, th, thead, tr} from './lib/html.js';
import {svgData, circle, defs, g, line, path, pattern, rect, svg, text, use} from './lib/svg.js';
import {windows, desktop, shell} from './lib/windows.js';
import {data, users, alarms, lines, reasons} from './data.js';

declare const pageLoad: Promise<void>;

let minuteWidth = 20,
    nextLabelID = 0;
const userFilter = Array.from({"length": users.length}, () => true),
      lineHighlight = Array.from({"length": lines.length}, () => false),
      thresholds: [number, string][] = [
	      [30, "#0f0"],
	      [80, "#ff0"],
	      [240, "#fa0"],
	      [Infinity, "#f00"]
      ],
      ml = div({"id": "mouseLine"}),
      mlt = div({"id": "mouseLine2"}),
      mt = div({"style": {"background-color": "#fff", "border": "1px solid #000", "position": "absolute", "top": 0, "z-index": 10}}),
      timeline = div({"class": "sticky"}),
      s = shell(desktop(timeline)),
      pad = (n: number) => (n < 10 ? "0" : "") + n,
      formatTime = (time: number) => {
	const d = new Date(time * 1000);
	return `${d.getFullYear()}-${pad(d.getMonth()+1)}-${d.getDate()} ${pad(d.getHours())}:${pad(d.getMinutes())}:${pad(d.getSeconds())}`;
      },
      formatDuration = (duration: number) => `${Math.floor(duration / 3600)}:${pad(Math.floor((duration % 3600) / 60))}:${pad(duration % 60)}`,
      addLabel = (name: string, input: HTMLInputElement): [HTMLLabelElement, HTMLInputElement] => {
	const id = "ID_" + nextLabelID++;
	return [label({"for": id}, name), amendNode(input, {id})];
      },
      stringSort = new Intl.Collator().compare,
      ss = (a: [string, number], b: [string, number]) => stringSort(a[0], b[0]),
      settingsButton = svg({"viewBox": "0 0 20 20", "onclick": () => {
	s.addWindow(settingsWindow);
	settingsWindow.focus();
      }}, [
	defs(path({"id": "spoke", "d": "M1,7 v2 a1,1 0,0,1 -2,0 v-2 z", "fill": "#aaa"})),
	g({"transform": "translate(10, 10)"}, [
		circle({"r": 5.5, "fill": "none", "stroke": "#aaa", "stroke-width": 4.5}),
		Array.from({"length": 12}, (_, n) => n * 30).map(r => use({"href": "#spoke", "transform": `rotate(${r})`}))
	])
      ]),
      settings = th({"rowspan": 2}, div(settingsButton)),
      settingsWindow = windows({"window-icon": svgData(settingsButton).replaceAll("aaa", "000"), "window-title": "Settings", "tabindex": -1, "onkeydown": (e: KeyboardEvent) => {
	if (e.key === "Escape") {
		settingsWindow.close();
	}
      }}, [
	h1("Settings"),
	addLabel("Timeline Scale (pixels per minute): ", input({"type": "number", "min": 1, "value": minuteWidth, "onchange": function(this: HTMLInputElement) {minuteWidth = parseInt(this.value);}})),
	br(),
	addLabel("Floating Logtime: ", input({"type": "checkbox", "checked": true, "onchange": function(this: HTMLInputElement) {amendNode(timeline, {"class": this.checked ? "sticky" : undefined});}})),
	br(),
	br(),
	table([
		tr([
			th(div({"style": {"text-decoration": "underline"}}, "Toggle Users:")),
			th(div({"style": {"text-decoration": "underline"}}, "Highlight Lines:"))
		]),
		tr([
			td((users.map((user, n) => [user, n]) as [string, number][]).sort(ss).map(([user, n]) => [
				addLabel(`${user}: `, input({"type": "checkbox", "checked": true, "onclick": function(this: HTMLInputElement) {userFilter[n] = this.checked}})),
				br()
			])),
			td((lines.map((line, n) => [line, n]) as [string, number][]).sort(ss).map(([line, n]) => [
				addLabel(`${line}: `, input({"type": "checkbox", "onclick": function(this: HTMLInputElement) {lineHighlight[n] = this.checked}})),
				br()
			])),
		]),
	]),
	button({"onclick": () => buildTimeline(data)}, "Rebuild")
      ]),
      buildTimeline = (data: Data) => {
	const rows = new Map<number, [HTMLDivElement, HTMLDivElement, Data[], [number, number], [number, number][]]>(),
	      tb = tbody(),
	      loggedRows: Data[] = [],
	      mm = (e: MouseEvent) => {
		const offset = e.offsetX + (e.target instanceof HTMLDivElement ? e.target.offsetLeft : 0);
		amendNode(ml, {"style": {"left": settings.offsetWidth + offset + "px"}});
		amendNode(mlt, {"style": {"left": offset + "px"}});
		clearNode(mt, {"style": {"left": settings.offsetWidth + offset + ""}}, formatTime(earliest + 60 * (offset + 2) / minuteWidth));
	      },
	      hours: SVGTextElement[] = [];
	let earliest = Infinity,
	    latest = -Infinity,
	    maxRows = 0;
	for (const row of data) {
		const [user, start, stop, logged] = row;
		if (!userFilter[user]) {
			continue;
		}
		let d: Data[],
		    set = false,
		    cc: [number, number],
		    dd: [number, number][];
		if (!rows.has(user)) {
			const h = th(div(users[user])),
			      t = td({"onmousemove": mm});
			rows.set(user, [h, t, d = [], cc = [-1, 0], dd = []]);
			amendNode(tb, tr([h, t]));
		} else {
			[, , d, cc, dd] = rows.get(user)!;
		}
		if (start > cc[0]) {
			if (cc[0] > -1) {
				dd.push([cc[0], start]);
			}
			cc[1] += stop - start;
			cc[0] = stop;
		} else if (stop > cc[0]) {
			cc[1] += stop - cc[0];
			cc[0] = stop;
		}
		DLoop:
		for (const r of d) {
			for (const [, cstart, cstop] of r) {
				if (start < cstop && stop > cstart) {
					continue DLoop;
				}
			}
			set = true;
			r.push(row);
			break;
		}
		if (!set) {
			d.push([row])
		}
		set = false;
		if (logged > 0) {
			LLoop:
			for (const r of loggedRows) {
				for (const [, cstop,, cstart] of r) {
					if (logged < cstop && start > cstart) {
						continue LLoop;
					}
				}
				set = true;
				r.push(row);
				break;
			}
			if (!set) {
				loggedRows.push([row]);
			}
			if (logged < earliest) {
				earliest = logged;
			}
		}
		if (start < earliest) {
			earliest = start;
		}
		if (stop > latest) {
			latest = stop;
		}
	}
	earliest = Math.floor(earliest / 60) * 60;
	latest = Math.ceil(latest / 60) * 60;
	for (const [, [t, cell, d, [, cc], dd]] of rows) {
		let rnum = 0,
		    calls = 0,
		    secs = 0;
		for (const [start, end] of dd) {
			amendNode(cell, div({"class": "wait", "title": `Down Time: ${formatDuration(end - start)}`, "style": {"left": (minuteWidth * (start - earliest) / 60 - 2) + "px", "width": (minuteWidth * (end - start) / 60 + 1) + "px"}}));
		}
		for (const row of d) {
			for (const [, start, stop,, line, reason, alarm] of row) {
				amendNode(cell, div({"class": lineHighlight[line] ? "h" : undefined, "title": `${formatTime(start)} ⟶  ${formatTime(stop)}\nCall Time: ${formatDuration(stop - start)}${alarm ? `\n${alarms[alarm]}` : ""}\n${reasons[reason]}`, "style": {"width": (minuteWidth * (stop - start) / 60 + 1) + "px", "left": (minuteWidth * (start - earliest) / 60 - 2) + "px", "--row": rnum}}));
				calls++;
				secs += stop - start;
			}
			rnum++;
		}
		amendNode(t, {"title": `Total Calls: ${calls}\nTotal Call Time: ${formatDuration(secs)}\nNon-concurrent Call Time: ${formatDuration(cc)}`});
		maxRows = Math.max(d.length, maxRows);
	}
	amendNode(ml, {"style": {"--h": maxRows * rows.size + loggedRows.length}});
	amendNode(document.body, {"style": {"--rows": maxRows}});
	for (let hour = Math.ceil(earliest / 3600) * 3600; hour <= latest; hour += 3600) {
		const d = new Date(hour * 1000);
		hours.push(text({"x": (minuteWidth * (hour - earliest) / 60), "y": 15}, `${pad(d.getHours())}:00`));
	}
	amendNode(mt, {"style": {"left": "-1000px"}});
	clearNode(timeline, table([
		thead({"style": {"--rows": loggedRows.length}}, [
			tr([settings, td({"style": {"width": (minuteWidth * (latest - earliest) / 60) + "px"}}, [
				svg({"width": minuteWidth * (latest - earliest) / 60, "height": 20, "viewBox": `0 0 ${minuteWidth * (latest - earliest) / 60} 20`}, [
					defs(pattern({"id": "ticks", "patternUnits": "userSpaceOnUse", "width": 60 * minuteWidth, "height": 20, "x": -(earliest % 3600) / 60 * minuteWidth - 2}, [
						line({"y2": 20, "stroke": "#000"}),
						line({"x1": minuteWidth * 15, "x2": minuteWidth * 15, "y2": 10, "stroke": "#000"}),
						line({"x1": minuteWidth * 30, "x2": minuteWidth * 30, "y2": 10, "stroke": "#000"}),
						line({"x1": minuteWidth * 45, "x2": minuteWidth * 45, "y2": 10, "stroke": "#000"})
					])),
					rect({"width": "100%", "height": "100%", "fill": "url(#ticks)"}),
					hours
				]),
				mt
			])]),
			tr(td({"onmousemove": mm}, [
				mlt,
				loggedRows.map((row, n) => row.map(([user, start,, logged, line,, alarm]) => div({"class": lineHighlight[line] ? "h" : undefined, "title": `${users[user]}\n${formatTime(logged)} ⟶  ${formatTime(start)}\nWait Time: ${formatDuration(start - logged)}${alarm ? `\n${alarms[alarm]}` : ""}`, "style": {"width": (minuteWidth * (start - logged) / 60 + 1) + "px", "left": (minuteWidth * (logged - earliest) / 60 - 2) + "px" , "--row": n, "color": thresholds.find(([max]) => max > (start - logged))![1]}})))
			]))
		]),
		tb
	]));
      };

data.sort(([, a], [, b]) => a - b);

pageLoad.then(() => {
	clearNode(document.body, [
		s,
		ml
	]);
	buildTimeline(data);
});
