const sliders = document.querySelectorAll('.slider-cpu');

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

function processUsageStats(usageStats) {
	sliders.forEach((current, index) => {
		const value = usageStats[index].toFixed(0);
		const percentageValue = value + '%';
		changeColorByValue(current, value);
		current.style.width = percentageValue;
		current.innerText = percentageValue;
		current.setAttribute('aria-valuenow', value);
	});
}

function updateSliders() {
	fetch(`${location.origin}/cpuusage`)
	.then(body => body.json())
	.then(usageStats => processUsageStats(usageStats));
}

setInterval(updateSliders, 500);

const startButton = document.querySelector('#start-button');
const magnitudeField = document.querySelector('#magnitude');
const unitField = document.querySelector('#unit');

startButton.addEventListener('click', e => {
	e.preventDefault();
	fetch('/cpustress', {
		method: 'POST',
		body: new URLSearchParams(`magnitude=${magnitudeField.value}&unit=${unitField.value}`),
	}).catch(console.err);
});
