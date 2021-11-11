import React, { Component } from 'react'
import PieChart from './PieChart'
//import StackedColumnChart from './StackedColumnChart'
//import GenericReport from './GenericReport'
//import './cards.css'
//import './card.css'

//const API = "https://20211002t224125-dot-vibrant-tree-324821.appspot.com/api/tweetscount";
//const API2 = "https://20211002t224125-dot-vibrant-tree-324821.appspot.com/api/distincthashtagscount";
//const API3 = "https://20211002t224125-dot-vibrant-tree-324821.appspot.com/api/upvotescount";

export class Reports extends Component {
    /*state = {
        tweetsCount: 0,
        differentHashtagsCount: 0,
        totalUpvotesCount: 0
    }

    componentDidMount() {
        fetch(`${API}`)
        .then((response) => response.json())
        .then(tweetCounter => {
            console.log(tweetCounter[0].CantidadTweets)
            this.setState({ tweetsCount: tweetCounter[0].CantidadTweets });
            
        });

        fetch(`${API2}`)
        .then((response) => response.json())
        .then(differentHashCount => {
            this.setState({ differentHashtagsCount: differentHashCount[0].DistinctHashtagsCount });
        });

        fetch(`${API3}`)
        .then((response) => response.json())
        .then(upvotesCounter => {
            this.setState({ totalUpvotesCount: upvotesCounter[0].UpvotesCount });
        });
    }
*/
    render() {
        return (
            <div className="container d-flex flex-column justify-content-center align-items-center">
                <div className="row">
                    <div className="col-md-6" style={{width: 600, padding:5}}>
                        <PieChart/>
                    </div>
                    <div className="col-md-6" style={{width: 600, padding:5}}>
                        
                    </div>
                </div>
            </div>
        )
    }
}

export default Reports
