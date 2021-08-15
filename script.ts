import type {Data} from './data.js';
import {createHTML, clearElement} from './lib/dom.js';
import {div, table, tbody, td, th, thead, tr,} from './lib/html.js';
import {defs, line, pattern, rect, svg, text} from './lib/svg.js';
import {data} from './data.js';

declare const pageLoad: Promise<void>;

const minuteWidth = 10,
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
      timeline = div(),
      ml = div({"id": "mouseLine"}),
      mm = (e: MouseEvent) => ml.style.setProperty("--x", e.offsetX + (e.target instanceof HTMLDivElement ? e.target.offsetLeft : 0) + "px"),
      pad = (n: number) => {
	const t = n + "";
	return t.length == 1 ? "0" + t : t;
      },
      formatTime = (time: number) => {
	const d = new Date(time * 1000);
	return `${d.getFullYear()}-${pad(d.getMonth()+1)}-${d.getDate()} ${pad(d.getHours())}:${pad(d.getMinutes())}:${pad(d.getSeconds())}`;
      },
      roundTimes = (start: number, stop: number) => [Math.floor(start / 60) * 60, Math.ceil(stop / 60) * 60],
      buildTimeline = (data: Data) => {
	const rows = new Map<string, [HTMLDivElement, HTMLDivElement, Data[]]>(),
	      tb = tbody();
	let earliest = Infinity,
	    latest = -Infinity,
	    numRows = 0;
	for (const row of data) {
		const [user, start, stop] = row;
		let d: [string, number, number][][],
		    set = false;
		if (!rows.has(user)) {
			const h = div(user),
			      t = td({"style": {"color": colours[rows.size % colours.length]}, "onmousemove": mm});
			rows.set(user, [h, t, d = []]);
			tb.appendChild(tr([th(h), t]));
		} else {
			[, , d] = rows.get(user)!;
		}
		Loop:
		for (const r of d) {
			for (const [, cstart, cstop] of r) {
				if (start < cstop && stop > cstart) {
					continue Loop
				}
			}
			set = true;
			r.push(row);
			break;
		}
		if (!set) {
			d.push([row])
			numRows++;
		}
		if (start < earliest) {
			earliest = start;
		}
		if (stop > latest) {
			latest = stop;
		}
	}
	ml.style.setProperty("--h",  numRows + "");
	[earliest, latest] = roundTimes(earliest, latest);
	for (const [, [t, cell, d]] of rows) {
		let rnum = 0;
		for (const row of d) {
			for (const [, start, stop] of row) {
				const [a, b] = roundTimes(start, stop);
				cell.appendChild(div({"title": formatTime(start) + " ⟶ " + formatTime(stop), "style": {"width": (minuteWidth * (b - a) / 60) + "px", "left": (minuteWidth * (a - earliest) / 60) + "px", "--row": rnum}}));
			}
			rnum++;
		}
		t.style.setProperty("--rows", rnum+"");
	}
	const hours: SVGTextElement[] = [];
	for (let hour = Math.ceil(earliest / 3600) * 3600; hour <= latest; hour += 3600) {
		const d = new Date(hour * 1000);
		hours.push(text({"x": (minuteWidth * (hour - earliest) / 60), "y": 15}, `${pad(d.getHours())}:00`));
	}
	createHTML(clearElement(timeline), table([
		thead(tr([td(), th({"style": {"width": (minuteWidth * (latest - earliest) / 60) + "px"}}, svg({"width": minuteWidth * (latest - earliest) / 60, "height": 20, "viewBox": `0 0 ${minuteWidth * (latest - earliest) / 60} 20`}, [
			defs(pattern({"id": "ticks", "patternUnits": "userSpaceOnUse", "width": 60 * minuteWidth, "height": 20, "x": -(earliest % 3600) / 60 * minuteWidth}, [
				line({"y2": 20, "stroke": "#000"}),
				line({"x1": minuteWidth * 15, "x2": minuteWidth * 15, "y2": 10, "stroke": "#000"}),
				line({"x1": minuteWidth * 30, "x2": minuteWidth * 30, "y2": 10, "stroke": "#000"}),
				line({"x1": minuteWidth * 45, "x2": minuteWidth * 45, "y2": 10, "stroke": "#000"})
			])),
			rect({"width": "100%", "height": "100%", "fill": "url(#ticks)"}),
			hours
		]))])),
		tb,
	]));
      };

pageLoad.then(() => {
	createHTML(clearElement(document.body), [
		div(),
		timeline,
		ml
	]);
	buildTimeline(data);
});
