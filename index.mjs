import { h, render } from "https://unpkg.com/preact?module";
import htm from "https://unpkg.com/htm?module";

const html = htm.bind(h);

function MemoryUsageBar(props) {
  return html`
    <span>
      <div class="bar">
        <div class="bar-inner" style="width: ${props.memUsagePercent}%"></div>
        <label>
          ${props.memUsage}MiB/${props.memTotal}MiB (${props.memUsagePercent}%)
        </label>
      </div>
    </span>
  `;
}

function CpuUsageBar(props) {
  return html`
  <span>
    <span>
      <h4>Usage:</h4>
      ${props.cpus.map((cpu) => {
        return html`<div class="bar">
          <div class="bar-inner" style="width: ${cpu}%"></div>
          <label>${cpu.toFixed(2)}%</label>
        </div>`;
      })}
    </span>
    <span>
      <h4>Confidence Interval:</h4>
       <span class="flex-items margin30">
        <span>
        <h4>90s:</h4>
        <div class="bar">
          <div class="bar-inner" style="width: ${props.cInterval90s[0]}%"></div>
          <label>Lowest: ${props.cInterval90s[0]}%</label>
        </div>
                <div class="bar">
          <div class="bar-inner" style="width: ${props.cInterval90s[1]}%"></div>
          <label>Highest: ${props.cInterval90s[1]}%</label>
        </div>
          <div class="bar">
          <div class="bar-inner" style="width: ${props.usageAverage90s}%"></div>
          <label>Average: ${props.usageAverage90s}%</label>
          </div>
      </span>
      <span>
        <h4>900s:</h4>
        <div class="bar">
          <div class="bar-inner" style="width: ${
            props.cInterval900s[0]
          }%"></div>
          <label>Lowest: ${props.cInterval900s[0]}%</label>
          </div>
          <div class="bar">
          <div class="bar-inner" style="width: ${
            props.cInterval900s[1]
          }%"></div>
          <label>Highest: ${props.cInterval900s[1]}%</label>
          </div>
          <div class="bar">
          <div class="bar-inner" style="width: ${
            props.usageAverage900s
          }%"></div>
          <label>Average: ${props.usageAverage900s}%</label>
      </span>
    </span>
   </span>
  </span>
    </span>

  `;
}

function App(props) {
  return html`
    <span>
      <h1>Hi! ${props.hostname}</h1>
    </span>
    <span>
      <span>
        <h4>Kernel: ${props.kernelName}</h4>
        <h4>Distro Name: ${props.osName}</h4>
        <h4>Gpu: ${props.gpuName}</h4>
        <h4>Memory:</h4>
        <span id="memUsage"> </span>
        <h4>Cpu: ${props.cpuName}</h4>
      </span>
      <span id="cpuUsage"> </span>
    </span>
  `;
}
let memTotal = await (await fetch("/system/mem/total")).json();
let init = async () => {
  let hostname = await (await fetch("/system/os/hostname")).json();
  let osName = await (await fetch("/system/os/name")).json();
  let kernelName = await (await fetch("/system/os/kernel")).json();
  let cpuName = await (await fetch("/system/cpu/name")).json();
  let gpuName = await (await fetch("/system/gpu/name")).json();
  // let memAvailable = await (await fetch("/system/mem/available")).json();
  let memUsage = await (await fetch("/system/mem/usage")).json();
  let memUsagePercent = await (await fetch("/system/mem/usagepercent")).json();

  render(
    html`<${App} hostname=${hostname} kernelName=${kernelName} osName=${osName} cpuName=${cpuName} gpuName=${gpuName} memTotal=${memTotal} memUsagePercent=${memUsagePercent} memUsage=${memUsage}></${App}>`,
    document.body
  );
};
let updateCpu = async () => {
  let cpus = await (await fetch("/realtime/cpus")).json();
  let cInterval90s = await (await fetch("/realtime/cpus/cinterval/90")).json();
  let usageAverage90s = await (await fetch("/realtime/cpus/average/90")).json();
  let cInterval900s = await (
    await fetch("/realtime/cpus/cinterval/900")
  ).json();
  let usageAverage900s = await (
    await fetch("/realtime/cpus/average/900")
  ).json();

  render(
    html`<${CpuUsageBar} cpus=${cpus} cInterval90s=${cInterval90s} usageAverage90s=${usageAverage90s} cInterval900s=${cInterval900s} usageAverage900s=${usageAverage900s}></${App}>`,
    document.getElementById("cpuUsage")
  );
};

let updateMemory = async () => {
  let memUsage = await (await fetch("/system/mem/usage")).json();
  let memUsagePercent = await (await fetch("/system/mem/usagepercent")).json();
  render(
    html`<${MemoryUsageBar} memUsage=${memUsage} memTotal=${memTotal} memUsagePercent=${memUsagePercent}></${App}>`,
    document.getElementById("memUsage")
  );
};

init();
updateCpu();
updateMemory();
setInterval(updateCpu, 500);
setInterval(updateMemory, 1000);
