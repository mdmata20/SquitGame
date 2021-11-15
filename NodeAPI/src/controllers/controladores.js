//const mysqlConnection = require('../conexion');

const MongoClient = require('mongodb').MongoClient;
const assert = require('assert');

// Connection URL
const url = 'mongodb://34.125.189.71:27017';

// Database Name
const dbName = 'squid-game';

const client = new MongoClient(url);

const findGamesCount = function(db, callback) {
    // Get the documents collection
    const collection = db.collection('games');
    // Find some documents
    collection.countDocuments({},function(err, docs) {
      assert.equal(err, null);
      console.log("Found the following records");
      console.log(docs)
      callback(docs);
    });
}

const findGames = function(db, callback) {
    // Get the documents collection
    const collection = db.collection('games');
    // Find some documents
    collection.find({}).sort({_id:-1}).toArray(function(err, docs) {
      assert.equal(err, null);
      console.log("Found the following records");
      console.log(docs)
      callback(docs);
    });
}

const findTop10 = function(db, callback) {
    // Get the documents collection
    const collection = db.collection('games');
    // Find some documents
    collection.find().sort({_id:-1}).limit(10).toArray(function(err, games) {
      assert.equal(err, null);
      console.log("Found the following records");
      console.log(games)
      callback(games);
    });
}

const findBest10Players = function(db, callback) {
    // Get the documents collection
    const collection = db.collection('games');
    // Find some documents
    collection.aggregate([{'$group': {_id: '$max', count:{$sum:1}}},{$sort:{"count":-1}}]).limit(10).toArray(function(err, games) {
      assert.equal(err, null);
      console.log("Found the following records");
      console.log(games)
      callback(games);
    });
}

const findGamesByWinner = function(db, callback, id) {
    // Get the documents collection
    const collection = db.collection('games');
    // Find some documents
    collection.find({"max":Number(id)}).toArray(function(err, docs) {
      assert.equal(err, null);
      console.log("Found the following records");
      console.log(id);
      console.log(docs)
      callback(docs);
    });
}

const findTop3Games = function(db, callback) {
    // Get the documents collection
    const collection = db.collection('games');
    // Find some documents
    collection.aggregate([{'$group': {_id: '$juego', y:{$sum:1}}},{$sort:{"y":-1}},{$limit:3}]).toArray(function(err, games) {
      assert.equal(err, null);
      console.log("Found the following records");
      console.log(games)
      callback(games);
    });
}

const findWorkers = function(db, callback) {
    // Get the documents collection
    const collection = db.collection('games');
    // Find some documents
    collection.aggregate([{'$group': {_id: '$worker', y:{$sum:1}}},{$sort:{"y":-1}},{$limit:3}]).toArray(function(err, games) {
      assert.equal(err, null);
      console.log("Found the following records");
      console.log(games)
      callback(games);
    });
}

const GetAllGamesCount = (req, res) => {
    // Use connect method to connect to the server
    client.connect(function(err) {
        assert.equal(null, err);
        console.log("Connected correctly to server");
    
        const db = client.db(dbName);
    
        findGamesCount(db, function(docs) {
            client.close();
            res.json(docs);
        });
    });
}

const GetAllGames = (req, res) => {
    // Use connect method to connect to the server
    client.connect(function(err) {
        assert.equal(null, err);
        console.log("Connected correctly to server");
    
        const db = client.db(dbName);
    
        findGames(db, function(docs) {
            client.close();
            res.json(docs);
        });
    });
}

const GetLast10Games = (req, res) => {
    // Use connect method to connect to the server
    client.connect(function(err) {
        assert.equal(null, err);
        console.log("Connected correctly to server");
    
        const db = client.db(dbName);
    
        findTop10(db, function(games) {
            client.close();
            res.json(games);
        });
    });
}

const GetBest10Players = (req, res) => {
    // Use connect method to connect to the server
    client.connect(function(err) {
        assert.equal(null, err);
        console.log("Connected correctly to server");
    
        const db = client.db(dbName);
    
        findBest10Players(db, function(games) {
            client.close();
            res.json(games);
        });
    });
}

const GetAllGamesByWinner = (req, res) => {

    var id = req.params.winner;

    // Use connect method to connect to the server
    client.connect(function(err) {
        assert.equal(null, err);
        console.log("Connected correctly to server");
    
        const db = client.db(dbName);
    
        findGamesByWinner(db, function(docs) {
            client.close();
            res.json(docs);
        }, id);
    });
}

const GetTop3Games = (req, res) => {
    // Use connect method to connect to the server
    client.connect(function(err) {
        assert.equal(null, err);
        console.log("Connected correctly to server");
    
        const db = client.db(dbName);
    
        findTop3Games(db, function(games) {
            client.close();
            res.json(games);
        });
    });
}

const GetWorkers = (req, res) => {
    // Use connect method to connect to the server
    client.connect(function(err) {
        assert.equal(null, err);
        console.log("Connected correctly to server");
    
        const db = client.db(dbName);
    
        findWorkers(db, function(games) {
            client.close();
            res.json(games);
        });
    });
}

module.exports = {
    GetAllGames : GetAllGames,
    GetLast10Games : GetLast10Games,
    GetBest10Players : GetBest10Players,
    GetAllGamesByWinner : GetAllGamesByWinner,
    GetTop3Games : GetTop3Games,
    GetWorkers : GetWorkers,
    GetAllGamesCount : GetAllGamesCount
}