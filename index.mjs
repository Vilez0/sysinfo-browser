import { h, render } from "https://unpkg.com/preact?module";
import htm from "https://unpkg.com/htm?module";

const html = htm.bind(h);

init();
async function init() {
  renderHostname();
  renderDistroName();
  renderKernelName();
  renderDesktop();
  renderCpuName();
  renderGpuName();
  renderUptime();
  renderDisks();
  updateCpu();
  updateMemory();
  updatecInterval(90);
  updatecInterval(900);
  setInterval(updateCpu, 1000);
  setInterval(updateMemory, 2500);
  setInterval(updatecInterval(90), 5000);
  setInterval(updatecInterval(900), 15000);
}

// This code renders a component to the DOM based on the element ID that is passed in
function renderer(elementId, component) {
  render(html`${component}`, document.getElementById(elementId));
}

// Thus functions retrieve info from server and render it to the DOM
async function renderHostname() {
  await fetch("/info/os/hostname")
    .then((res) => res.json())
    .then((res) => renderer("Hostname", res.data));
}

async function renderHostname() {
  await fetch("/info/os/hostname")
    .then((res) => res.json())
    .then((res) => renderer("Hostname", res.data));
}

async function renderHostname() {
  await fetch("/info/os/hostname")
    .then((res) => res.json())
    .then((res) => renderer("Hostname", res.data));
}

async function renderDistroName() {
  await fetch("/info/os/name")
    .then((res) => res.json())
    .then((res) => renderer("DistroName", res.data));
}

async function renderKernelName() {
  await fetch("/info/os/kernel")
    .then((res) => res.json())
    .then((res) => renderer("KernelName", res.data));
}

async function renderDesktop() {
  await fetch("/info/os/desktop")
    .then((res) => res.json())
    .then((res) => renderer("Desktop", res.data));
}

async function renderCpuName() {
  await fetch("/info/cpu/name")
    .then((res) => res.json())
    .then((res) => renderer("CpuName", res.data));
}

async function renderGpuName() {
  await fetch("/info/gpu/name")
    .then((res) => res.json())
    .then((res) => renderer("GpuName", res.data));
}

async function renderUptime() {
  await fetch("/info/os/uptime")
    .then((res) => res.json())
    .then((res) => renderer("Uptime", res.data));
}

async function renderDisks() {
  let disks = await fetch("/info/disks").then((res) => res.json());
  await Promise.all(
    disks.data.map(async (disk) => {
      let size = await fetch(`/info/disks/${disk}/size`).then((res) =>
        res.json()
      );
      document.getElementById(
        "diskUsage"
      ).innerHTML += `<h4>${disk} ${size.data}GiB</h4>`;
    })
  );
}

async function updateCpu() {
  let cpus = await fetch("/info/cpu/usage").then((res) => res.json());
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
  let cIntervals = await fetch(`/info/cpu/usage/cinterval/${seconds}`).then(
    (res) => res.json()
  );
  let usageAverage = await fetch(`/info/cpu/usage/average/${seconds}`).then(
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
  let memUsage = await fetch("/info/mem/usage").then((res) => res.json());
  let memUsagePercent = await fetch("/info/mem/usagepercent").then((res) =>
    res.json()
  );
  let memTotal = await fetch("/info/mem/total").then((res) => res.json());
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
