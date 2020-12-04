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
	});
}

function updateSliders() {
	fetch(`${location.origin}/cpuusage`)
	.then(body => body.json())
	.then(usageStats => processUsageStats(usageStats));
}

setInterval(updateSliders, 600);
