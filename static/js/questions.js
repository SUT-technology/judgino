const express = require('express');
const app = express();
const { Pool } = require('pg'); // Assuming you're using PostgreSQL

const pool = new Pool({
  user: 'your_username',
  host: 'localhost',
  database: 'your_database',
  password: 'your_password',
  port: 5432,
});

app.get('/questions', async (req, res) => {
    try {
        const result = await pool.query('SELECT title, publish_date, deadline FROM questions');
        res.json(result.rows); // Sending the questions data as a JSON response
    } catch (err) {
        console.error('Error executing query', err.stack);
        res.status(500).send('Error fetching questions');
    }
});

app.listen(3000, () => {
    console.log('Server running on http://localhost:3000');
});