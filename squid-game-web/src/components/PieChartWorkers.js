import React, { Component } from 'react';
import CanvasJSReact from '../assets/canvasjs.react';
import socketIOClient from "socket.io-client";

var CanvasJSChart = CanvasJSReact.CanvasJSChart;

const API2 = "http://35.225.182.66:4001/api/games/workers";
const ENDPOINT2 = "http://35.225.182.66:4001";
 
class PieChartWorkers extends Component {
    state = {
        workers: []
    }

    componentDidMount() {
        fetch(`${API2}`)
        .then((response) => response.json())
        .then(workersStats => {
            this.setState({ workers: workersStats });
            console.log(this.state.workers);
        });

		const socket = socketIOClient(ENDPOINT2);
      
        socket.on("NewGamesNotify", data => {
            console.log('the db has changed!')
            this.fetchWorkers();
        //toast.info('New Tweets have been pushed!', 
        //{position: toast.POSITION.TOP_CENTER});
        });
    }

	fetchWorkers = () => {
        fetch(`${API2}`)
        .then((response) => response.json())
        .then(workersStats => {
            this.setState({ workers: workersStats });
        });
    }

	render() {
		const options = {
			exportEnabled: true,
			animationEnabled: true,
			title: {
				text: "Workers Stats"
			},
			data: [{
				type: "pie",
				startAngle: 75,
				toolTipContent: "<b>{_id}</b>: {y} insertion(s)",
				showInLegend: "true",
				legendText: "{_id}",
				indexLabelFontSize: 16,
				indexLabel: "{_id} - {y}",
				dataPoints: this.state.workers
			}]
		}
		
		return (
			<div className="container d-flex flex-column justify-content-center align-items-center">
				<div className="row">
					<div className="col-md-6" style={{width: 600, padding:5}}>
						<div className="card">
							<div className="card-body">
								<CanvasJSChart options = {options} 
									/* onRef={ref => this.chart = ref} */
								/>
            				</div>
        				</div>
					</div>
        		</div>
			</div>
		);
	}
}

export default PieChartWorkers;