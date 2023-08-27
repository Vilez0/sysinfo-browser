import { h, render } from "https://unpkg.com/preact?module";
import htm from "https://unpkg.com/htm?module";

const html = htm.bind(h);

function renderer(elementId, component) {
  render(html`${component}`, document.getElementById(elementId));
}

async function init() {
  await fetch("/system/os/hostname")
    .then((res) => res.json())
    .then((res) => {
      renderer("Hostname", res.data);
    });

  await fetch("/system/os/hostname")
    .then((res) => res.json())
    .then((res) => renderer("Hostname", res.data));

  await fetch("/system/os/name")
    .then((res) => res.json())
    .then((res) => renderer("DistroName", res.data));
  await fetch("/system/os/kernel")
    .then((res) => res.json())
    .then((res) => renderer("KernelName", res.data));
  await fetch("/system/os/desktop")
    .then((res) => res.json())
    .then((res) => renderer("Desktop", res.data));
  await fetch("/system/cpu/name")
    .then((res) => res.json())
    .then((res) => renderer("CpuName", res.data));

  await fetch("/system/gpu/name")
    .then((res) => res.json())
    .then((res) => renderer("GpuName", res.data));
  renderDisks();
}

async function renderDisks() {
  let disks = await fetch("/system/disks").then((res) => res.json());
  await Promise.all(
    disks.data.map(async (disk) => {
      let size = await fetch(`/system/disks/${disk}/size`).then((res) =>
        res.json()
      );
      document.getElementById(
        "diskUsage"
      ).innerHTML += `<h4>${disk} ${size.data}GiB</h4>`;
    })
  );
}

async function updateCpu() {
  let cpus = await fetch("/system/cpu/usage").then((res) => res.json());
  renderer(
    "cpuUsage",
    html`${cpus.data.map((cpu) => {
      return html`<div class="bar">
        <div class="bar-inner" style="width: ${cpu}%"></div>
        <label>${cpu.toFixed(2)}%</label>
      </div>`;
    })}`
  );
}

async function updatecInterval(seconds) {
  let cIntervals = await fetch(`/system/cpu/usage/cinterval/${seconds}`).then(
    (res) => res.json()
  );
  let usageAverage = await fetch(`/system/cpu/usage/average/${seconds}`).then(
    (res) => res.json()
  );
  renderer(
    `c${seconds}s`,
    html`<div class="bar">
        <div class="bar-inner" style="width: ${cIntervals.data[0]}%"></div>
        <label>Lowest: ${cIntervals.data[0]}%</label>
      </div>
      <div class="bar">
        <div class="bar-inner" style="width: ${cIntervals.data[1]}%"></div>
        <label>Highest: ${cIntervals.data[1]}%</label>
      </div>
      <div class="bar">
        <div class="bar-inner" style="width: ${usageAverage.data}%"></div>
        <label>Average: ${usageAverage.data}%</label>
      </div>`
  );
}

async function updateMemory() {
  let memUsage = await fetch("/system/mem/usage").then((res) => res.json());
  let memUsagePercent = await fetch("/system/mem/usagepercent").then((res) =>
    res.json()
  );
  let memTotal = await fetch("/system/mem/total").then((res) => res.json());
  renderer(
    `memUsage`,
    html`
      <div class="bar">
        <div class="bar-inner" style="width: ${memUsagePercent.data}%"></div>
        <label>
          ${memUsage.data}MiB/${memTotal.data}MiB (${memUsagePercent.data}%)
        </label>
      </div>
    `
  );
}

init();
updateCpu();
updateMemory();
updatecInterval(90);
updatecInterval(900);
setInterval(updateCpu, 1000);
setInterval(updateMemory, 2500);
setInterval(updatecInterval(90), 5000);
setInterval(updatecInterval(900), 15000);
