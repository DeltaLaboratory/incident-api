package org.openapitools.client.apis

import org.openapitools.client.infrastructure.CollectionFormats.*
import retrofit2.http.*
import retrofit2.Call
import okhttp3.RequestBody
import com.squareup.moshi.Json

import org.openapitools.client.models.CommonErrorResponse
import org.openapitools.client.models.IncidentDeleteReportRequest
import org.openapitools.client.models.IncidentHelpRequest
import org.openapitools.client.models.IncidentQueryHelpRequest
import org.openapitools.client.models.IncidentQueryHelpResponse
import org.openapitools.client.models.IncidentQueryRequest
import org.openapitools.client.models.IncidentQueryResponse
import org.openapitools.client.models.IncidentReportRequest
import org.openapitools.client.models.IncidentReportResponse
import org.openapitools.client.models.IncidentVoteRequest

interface IncidentApi {
    /**
     * Help
     * Help
     * Responses:
     *  - 201: Created
     *  - 400: Bad Request
     *  - 500: Internal Server Error
     *
     * @param request Help request
     * @return [Call]<[Unit]>
     */
    @POST("incident/help")
    fun incidentHelpPost(@Body request: IncidentHelpRequest): Call<Unit>

    /**
     * Query help
     * Query help
     * Responses:
     *  - 200: OK
     *  - 400: Bad Request
     *  - 500: Internal Server Error
     *
     * @param request Query help request
     * @return [Call]<[kotlin.collections.List<IncidentQueryHelpResponse>]>
     */
    @POST("incident/help/query")
    fun incidentHelpQueryPost(@Body request: IncidentQueryHelpRequest): Call<kotlin.collections.List<IncidentQueryHelpResponse>>

    /**
     * Query incident image
     * Query incident image
     * Responses:
     *  - 200: OK
     *  - 400: Bad Request
     *  - 404: Not Found
     *  - 500: Internal Server Error
     *
     * @param incidentId Incident ID
     * @return [Call]<[kotlin.String]>
     */
    @GET("incident/image/{incident_id}")
    fun incidentImageIncidentIdGet(@Path("incident_id") incidentId: kotlin.String): Call<kotlin.String>

    /**
     * Query incidents
     * Query incidents
     * Responses:
     *  - 200: OK
     *  - 400: Bad Request
     *  - 500: Internal Server Error
     *
     * @param request Query request
     * @return [Call]<[kotlin.collections.List<IncidentQueryResponse>]>
     */
    @POST("incident/query")
    fun incidentQueryPost(@Body request: IncidentQueryRequest): Call<kotlin.collections.List<IncidentQueryResponse>>

    /**
     * Delete an incident report
     * Delete an incident report
     * Responses:
     *  - 204: No Content
     *  - 400: Bad Request
     *  - 404: Not Found
     *  - 500: Internal Server Error
     *
     * @param request Delete report request
     * @return [Call]<[Unit]>
     */
    @DELETE("incident/report")
    fun incidentReportDelete(@Body request: IncidentDeleteReportRequest): Call<Unit>

    /**
     * Report an incident
     * Report an incident
     * Responses:
     *  - 200: OK
     *  - 400: Bad Request
     *  - 500: Internal Server Error
     *
     * @param request Report request
     * @return [Call]<[IncidentReportResponse]>
     */
    @POST("incident/report")
    fun incidentReportPost(@Body request: IncidentReportRequest): Call<IncidentReportResponse>

    /**
     * Vote an incident
     * Vote an incident
     * Responses:
     *  - 204: No Content
     *  - 400: Bad Request
     *  - 404: Not Found
     *  - 409: Conflict
     *  - 500: Internal Server Error
     *
     * @param request Vote request
     * @return [Call]<[Unit]>
     */
    @POST("incident/vote")
    fun incidentVotePost(@Body request: IncidentVoteRequest): Call<Unit>

}
