const express       = require('express')
const bodyParser    = require('body-parser')

// create express app
const app = express()
app.use(bodyParser.json())

// register routes
app.post('/payment/charge', chargePayment)
app.get('/payment/charge', viewPaymentDetail)

// listen for requests
app.listen(8080, () => {
    console.log("Server is listening on port 8080")
})

async function chargePayment(req, res) {
    const body = { status: "Success" }
    res.status(200).send(body)
}

async function viewPaymentDetail(req, res) {
    const body = { 
        method: "CreditCard",
        status: "Success",
        creditCard: {
            number: "XXXXXXXXXXXX1111",
            holderName: "John Smith"
        }
     }
    res.status(200).send(body)
}