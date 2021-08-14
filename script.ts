import type {Data} from './data.js';
import {createHTML, clearElement} from './lib/dom.js';
import {div, table, tbody, td, th, thead, tr,} from './lib/html.js';
import {defs, line, pattern, rect, svg} from './lib/svg.js';
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
	const rows = new Map<string, HTMLDivElement>(),
	      tb = tbody();
	let earliest = Infinity,
	    latest = -Infinity;
	for (const [user, start, stop] of data) {
		if (!rows.has(user)) {
			const t = td({"style": {"color": colours[rows.size % colours.length]}});
			rows.set(user, t);
			tb.appendChild(tr([th(div(user)), t]));
		}
		if (start < earliest) {
			earliest = start;
		}
		if (stop > latest) {
			latest = stop;
		}
	}
	[earliest, latest] = roundTimes(earliest, latest);
	for (const [user, start, stop] of data) {
		const d = rows.get(user)!,
		      [a, b] = roundTimes(start, stop);
		d.appendChild(div({"title": formatTime(start) + " ⟶ " + formatTime(stop), "style": {"width": (minuteWidth * (b - a) / 60) + "px", "left": (minuteWidth * (a - earliest) / 60) + "px"}}));
	}
	createHTML(clearElement(timeline), table([
		thead(tr([td(), td({"style": {"width": (minuteWidth * (latest - earliest) / 60) + "px"}}, svg({"width": minuteWidth * (latest - earliest) / 60, "height": 20, "viewBox": `0 0 ${minuteWidth * (latest - earliest) / 60} 20`}, [
			defs(pattern({"id": "ticks", "patternUnits": "userSpaceOnUse", "width": 60 * minuteWidth, "height": 20}, [
				line({"y2": 20, "stroke": "#000"}),
				line({"x1": minuteWidth * 15, "x2": minuteWidth * 15, "y2": 10, "stroke": "#000"}),
				line({"x1": minuteWidth * 30, "x2": minuteWidth * 30, "y2": 10, "stroke": "#000"}),
				line({"x1": minuteWidth * 45, "x2": minuteWidth * 45, "y2": 10, "stroke": "#000"})
			])),
			rect({"width": "100%", "height": "100%", "fill": "url(#ticks)"}),
		]))])),
		tb,
	]));
      };

pageLoad.then(() => {
	createHTML(clearElement(document.body), [
		div(),
		timeline
	]);
	buildTimeline(data);
});
