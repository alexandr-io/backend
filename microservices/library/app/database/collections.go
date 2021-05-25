package database

import "go.mongodb.org/mongo-driver/mongo"

// CollectionBook is the name of the book collection in the database
const CollectionBook = "book"

// CollectionLibraries is the name of the user_library collection in the database
const CollectionLibraries = "user_library"

// CollectionLibrary is the name of the library collection in the database
const CollectionLibrary = "library"

// CollectionBookProgress is the name of the user's book progress collection in mongodb
const CollectionBookProgress = "book_progress"

// CollectionGroup is the name of the group collection in mongodb
const CollectionGroup = "group"

// CollectionUserData is the name of the user data collection in the database
const CollectionUserData = "user_data"

// BookCollection collection for book
var BookCollection *mongo.Collection

// LibrariesCollection collection for libraries
var LibrariesCollection *mongo.Collection

// LibraryCollection collection for library
var LibraryCollection *mongo.Collection

// BookProgressCollection collection for book progress
var BookProgressCollection *mongo.Collection

// GroupCollection collection for groups
var GroupCollection *mongo.Collection
