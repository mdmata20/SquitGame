import React, { Component } from 'react';
import CanvasJSReact from '../assets/canvasjs.react';
import socketIOClient from "socket.io-client";

var CanvasJSChart = CanvasJSReact.CanvasJSChart;

const API = "http://35.225.182.66:4001/api/games/top3games";
const ENDPOINT = "http://35.225.182.66:4001";
 
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

		const socket = socketIOClient(ENDPOINT);
      
        socket.on("NewGamesNotify", data => {
            console.log('the db has changed!')
            this.fetchTop3Games();
        //toast.info('New Tweets have been pushed!', 
        //{position: toast.POSITION.TOP_CENTER});
        });
    }

	fetchTop3Games = () => {
        fetch(`${API}`)
        .then((response) => response.json())
        .then(gamesList => {
            this.setState({ games: gamesList });
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

export default PieChart;