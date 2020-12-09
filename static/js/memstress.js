const slider = document.querySelector('#slider-memory');
const sizes = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
const k = 1024;

const initialValues = slider.innerText.split('/');
slider.innerText = formatBytes(initialValues[0]) + '/' + formatBytes(initialValues[1]);

function changeColorByValue(element, value) {
	let newClass = 'bg-success';
	if (value >= 50) {
		if (value >= 75) {
			newClass = 'bg-danger';
		} else {
			newClass = 'bg-warning';
		}
	}
	element.className = element.className.replace(/bg-(?:success|warning|danger)/, newClass);
}

function formatBytes(bytes, decimals = 2) {
	if (bytes === 0) return '0 Bytes';

	const dm = decimals < 0 ? 0 : decimals;

	const i = Math.floor(Math.log(bytes) / Math.log(k));

	return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
}

function processMemUsage(memUsage) {
	const percentageUsed = memUsage.percentage_used.toFixed(0);
	const percentageUsedString = percentageUsed + '%';
	const ratio = formatBytes(memUsage.used_memory) + '/' + formatBytes(memUsage.total_memory);
	slider.style.width = percentageUsedString;
	slider.innerText = ratio;
	slider.setAttribute('aria-valuenow', percentageUsed);
	changeColorByValue(slider, percentageUsed);
}

function updateSlider() {
	fetch(`${location.origin}/memusage`)
	.then(body => body.json())
	.then(memUsage => processMemUsage(memUsage));
}

setInterval(updateSlider, 500)

const startButton = document.querySelector('#start-button');
const magnitudeField = document.querySelector('#magnitude');
const unitField = document.querySelector('#unit');

startButton.addEventListener('click', e => {
	e.preventDefault();
	fetch('/memstress', {
		method: 'POST',
		body: new URLSearchParams(`magnitude=${magnitudeField.value}&unit=${unitField.value}`),
	}).catch(console.err);
});
