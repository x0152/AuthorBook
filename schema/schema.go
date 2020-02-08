package schema

var Schema = `
		schema {
			query: Query
		}
		# The query type, represents all of the entry points into our object graph
		type Query {
			book(id: ID!): Book 
			author(id: ID!): Author
			comment(id: ID!): Comment

			authors(): [Author!]!
			books(): [Book!]!
		}

		interface Book{
			id: ID!
			name: String!
			description: String!
			publishdate: String!
			author: Author!
		}

		interface Author{
			id: ID!
			name: String!
			secondname: String!
			books: [Book!]!
		}

		interface Comment{
			id: ID!
			text: String!
			book: Book!
			userid: ID!  
		}
	`
