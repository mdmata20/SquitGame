import React, { Component } from 'react';
import CanvasJSReact from '../assets/canvasjs.react';

var CanvasJSChart = CanvasJSReact.CanvasJSChart;

const API = "http://localhost:4001/api/games/top3games";
 
class PieChart extends Component {
    state = {
        games: []
    }

    componentDidMount() {
        fetch(`${API}`)
        .then((response) => response.json())
        .then(topGames => {
            this.setState({ games: topGames });
            console.log(this.state.games);
        });
    }

	render() {
		const options = {
			exportEnabled: true,
			animationEnabled: true,
			title: {
				text: "Top 3 Games"
			},
			data: [{
				type: "pie",
				startAngle: 75,
				toolTipContent: "<b>{_id}</b>: {y}",
				showInLegend: "true",
				legendText: "{_id}",
				indexLabelFontSize: 16,
				indexLabel: "{_id} - {y}",
				dataPoints: this.state.games
			}]
		}
		
		return (
		<div className="card">
            <div className="card-body">
			<CanvasJSChart options = {options} 
				/* onRef={ref => this.chart = ref} */
			/>
			{/*You can get reference to the chart instance as shown above using onRef. This allows you to access all chart properties and methods*/}
            </div>
        </div>
		);
	}
}

export default PieChart;