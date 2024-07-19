/**
 *
 * Please note:
 * This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * Do not edit this file manually.
 *
 */

@file:Suppress(
    "ArrayInDataClass",
    "EnumEntryName",
    "RemoveRedundantQualifierName",
    "UnusedImport"
)

package org.openapitools.client.models


import com.squareup.moshi.Json
import com.squareup.moshi.JsonClass

/**
 * 
 *
 * @param bloodPressure 
 * @param bodyTemperature 
 * @param createdAt 
 * @param description 
 * @param heartRate 
 * @param latitude 
 * @param longitude 
 * @param reporter 
 */


data class IncidentQueryHelpResponse (

    @Json(name = "blood_pressure")
    val bloodPressure: kotlin.Int? = null,

    @Json(name = "body_temperature")
    val bodyTemperature: kotlin.Int? = null,

    @Json(name = "created_at")
    val createdAt: kotlin.String? = null,

    @Json(name = "description")
    val description: kotlin.String? = null,

    @Json(name = "heart_rate")
    val heartRate: kotlin.Int? = null,

    @Json(name = "latitude")
    val latitude: java.math.BigDecimal? = null,

    @Json(name = "longitude")
    val longitude: java.math.BigDecimal? = null,

    @Json(name = "reporter")
    val reporter: kotlin.String? = null

) {


}

