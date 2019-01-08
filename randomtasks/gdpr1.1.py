import sys

from simtools.executable import make_this_loadable
from simtools.timezone import system_now, datetime
make_this_loadable()
import settings
from simtools.dbpool import make_pool
import multiprocessing
logger = multiprocessing.log_to_stderr()

pool = make_pool(settings)

blackjack = pool.get(pool.blackjack)
party = pool.get(pool.party)
prod = pool.get(pool.prod)
#subpub = pool.get(pool.subpub)

from boto.s3.connection import S3Connection
from boto.s3.connection import OrdinaryCallingFormat
S3_CONN = S3Connection(settings.AWS_ACCESS_KEY_ID, settings.AWS_SECRET_ACCESS_KEY)
S3_FIXED_CONN = S3Connection(settings.AWS_ACCESS_KEY_ID, settings.AWS_SECRET_ACCESS_KEY,
                             calling_format=OrdinaryCallingFormat())

S3_URL = "s3.amazonaws.com"
SELF_PROVIDER = "siminars.com"

KEYS = [
    "promotional_banner_url",
    "logo_url",
    "original_logo",
    "banner_url",
    "img_url",
    "css_url",
    "logo"
]

def _delete_remote(url):
    if not url:
        return

    parts = url.split(".s3.amazonaws.com")
    obj_path = parts[-1]
    bucket_name = parts[0].split("//")[-1]

    print "##### Before Deleting remote(url) #####"
    pprint(bucket_name)

    print "deleting %s from bucket %s" % (obj_path, bucket_name)
    # add code to delete
    try:
        logger.info("Connecting to: %s" % bucket_name)

        if "." in bucket_name:
            bucket = S3_FIXED_CONN.get_bucket(bucket_name)
        else:
            bucket = S3_CONN.get_bucket(bucket_name)

        bucket.delete_key(obj_path)
    except Exception, e:
        logger.error("Deletion FAILED: [%s] " % obj_path, exc_info=True)
        return

    print "##### After Deleting remote(url) #####"
    pprint(bucket_name)

    return

def _filter_remote_urls_from_assets(assets):
    urls = []
    for a in assets:
        if not a["provider_url"] or SELF_PROVIDER not in a["provider_url"]:
            continue

        _url = a.get("url")
        if _url and S3_URL in _url:
            urls.append(_url)

        author_logo = a["author"]["logo"]
        if author_logo and S3_URL in author_logo:
            urls.append(author_logo)

        thumbnail_url = a.get("thumbnail_url")
        if thumbnail_url and S3_URL in thumbnail_url:
            urls.append(thumbnail_url)

        _urls = a.get("urls")
        for url in _urls:
            if url and S3_URL in url:
                urls.append(url)

        streams = a.get("streams")
        for k in streams:
            if streams[k] and S3_URL in streams[k]:
                urls.append(streams[k])

    return list(set(urls))

def _filter_remote_urls(items):
    urls = []

    for i in items:
        if i.get("author"):
            author_logo = i["author"]["logo"]
            if author_logo and S3_URL in author_logo:
                urls.append(author_logo)

        for k in KEYS:
            url = i.get(k)
            if url and S3_URL in url:
                urls.append(url)
    return list(set(urls))

def _delete_remote_data(items, filter_fn):
    remote_urls = filter_fn(items)
    print "##### Before Deleting remote_data #####"
    pprint( remote_urls)
    for url in remote_urls:
        _delete_remote(url)
    print "##### After Deleting remote_data #####"
    pprint(remote_urls)


# Delete Blackjack Data
def delete_asset(uid):
    items = blackjack.asset.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting assets #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls_from_assets)

    blackjack.asset.update(
        {"author_id": uid},
        {"$set": {
            "provider_url": "deleted",
            "description": "deleted",
            "filetype": "deleted",
            "deleted": True,
            "author_name": "deleted",
            "thumbnail_url": None,
            "title": "deleted",
            "url": "deleted",
            "author.logo": None,
            "author.display_name": "Deleted",
            "html": None,
            "author_url": "deleted",
            "urls": [],
            "type": "deleted",
            "streams": {},
            "oembed": {}
        }},
        multi=True
    )

    items = blackjack.asset.find({"author_id": uid})
    items = [i for i in items]
    print "##### After Deleting assets #####"
    pprint(items)


def delete_bundle(uid):
    items = blackjack.bundle.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting bundle #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    blackjack.bundle.update(
        {"author_id": uid},
        {"$set": {
            "subtitle": "deleted",
            "topics": [],
            "featured_book": [],
            "promotional_banner_url": None,
            "title": "deleted",
            "logo_url": None,
            "blurb": "deleted",
            "description": "deleted",
            "deleted": True,
            "original_logo": None,
            "author.logo": None,
            "author.display_name": "Deleted"
        }},
        multi=True
    )

    items = blackjack.bundle.find({"author_id": uid})
    items = [i for i in items]
    print "##### After Deleting bundle #####"
    pprint(items)

def delete_community(uid):
    items = blackjack.community.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting community #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    blackjack.community.update(
        {"author_id": uid},
        {"$set": {
            "banner_url": None,
            "title": "deleted",
            "author.logo": None,
            "author.display_name": "Deleted",
            "deleted": True
        }},
        multi=True
    )
    print "##### After Deleting community #####"
    pprint(items)

def delete_coupon(uid):
    items = blackjack.coupon.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting coupon #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    blackjack.coupon.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "coupon": "deleted",
            "price": 0,
            "deleted": True
        }},
        multi=True
    )
    print "##### After Deleting coupon #####"
    pprint(items)

def delete_feedback(uid):
    items = blackjack.feedback.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting feedback #####"
    pprint(items)

    _delete_remote_data(items, _filter_remote_urls)

    blackjack.feedback.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "review_text": "deleted",
            "star_rating": "deleted",
            "deleted": True,
            "feedback_answers": []
        }},
        multi=True
    )
    print "##### After Deleting feedback #####"
    pprint(items)

def delete_jcomment(uid):
    items = blackjack.jcomment.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting jcomment #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    blackjack.jcomment.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "text": "deleted",
            "replies": [],
            "vote": {},
            "replies_count": 0
        }},
        multi=True
    )
    print "##### After Deleting jcomment #####"
    pprint(items)

def delete_jontent(uid):
    items = blackjack.jontent.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting jontent #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    blackjack.jontent.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "body": "deleted",
            "responses": [],
            "title": "deleted",
            "deleted": True
        }},
        multi=True
    )
    print "##### After Deleting jontent #####"
    pprint(items)

def delete_month(uid):
    items = blackjack.month.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting month #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    blackjack.month.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "description": "deleted",
            "title": "deleted"
        }},
        multi=True
    )
    print "##### After Deleting month #####"
    pprint(items)

def delete_quizanswer(uid):
    items = blackjack.quizanswer.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting quizanswer #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    blackjack.quizanswer.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "text": "deleted",
            "img_url": None,
            "is_correct": None
        }},
        multi=True
    )
    print "##### After Deleting quizanswer #####"
    pprint(items)

def delete_quizquestion(uid):
    items = blackjack.quizquestion.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting quizquestion #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    blackjack.quizquestion.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "text": "deleted",
            "answers": [],
            "correct_answers": [],
            "submission_text": "deleted",
            "method": None,
            "deleted": True
        }},
        multi=True
    )
    print "##### After Deleting quizquestion #####"
    pprint(items)

def delete_quizresult(uid):
    items = blackjack.quizresult.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting quizresult #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    blackjack.quizresult.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "total_score": 0,
            "question_answers": {},
            "score": 0,
            "questions": [],
            "score_range": {}
        }},
        multi=True
    )
    print "##### After Deleting quizresult #####"
    pprint(items)

def delete_response(uid):
    items = blackjack.response.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting response #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    blackjack.response.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "text": "deleted",
            "title": "deleted",
            "comments_count": 0
        }},
        multi=True
    )
    print "##### After Deleting response #####"
    pprint(items)

def delete_session(uid):
    items = blackjack.session.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting session #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    blackjack.session.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "subtitle": "deleted",
            "featured_book": "deleted",
            "list_price": 0,
            "promotional_banner_url": None,
            "title": "deleted",
            "timezone": None,
            "logo_url": None,
            "purchase_price": 0,
            "blurb": "deleted",
            "description": "deleted",
            "original_logo": None,
            "availability_time": None,
            "signupcode": None,
            "pacing": "deleted",
            "deleted": True
        }},
        multi=True
    )
    print "##### After Deleting session #####"
    pprint(items)

def delete_signup_code(uid):
    items = blackjack.signup_code.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting signup_code #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    blackjack.signup_code.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "code": "deleted",
            "use_count": None,
            "deleted": True,
            "use_limit": 0
        }},
        multi=True
    )
    print "##### After Deleting signup_code #####"
    pprint(items)

def delete_siminar(uid):
    items = blackjack.siminar.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting siminar #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    blackjack.siminar.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "subtitle": "deleted",
            "featured_book": "deleted",
            "list_price": 0,
            "promotional_banner_url": None,
            "banner_url": None,
            "title": "deleted",
            "logo_url": None,
            "blurb": "deleted",
            "description": "deleted",
            "original_logo": None,
            "confirmation_email": None
        }},
        multi=True
    )
    print "##### After Deleting siminar #####"
    pprint(items)

def delete_step(uid):
    items = blackjack.step.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting step #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    blackjack.step.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "description": "deleted",
            "title": "deleted"
        }},
        multi=True
    )
    print "##### After Deleting step #####"
    pprint(items)

def delete_subscribe(uid):
    items = blackjack.subscribe.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting subscribe #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    blackjack.siminar.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "subtitle": "deleted",
            "featured_book": None
        }},
        multi=True
    )
    print "##### After Deleting subscribe #####"
    pprint(items)

def delete_testimonial(uid):
    items = blackjack.testimonial.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting testimonial #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    blackjack.testimonial.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "source_name": "Deleted",
            "source_company": "Deleted",
            "text": "deleted"
        }},
        multi=True
    )
    print "##### After Deleting testimonial #####"
    pprint(items)


# prod_db
def delete_bill_plan(uid):
    items = prod.bill_plan.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting bill_plan #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    prod.bill_plan.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "title": "Deleted",
            "note": None,
            "logo_text": "Deleted",
            "cname": None,
            "email_sender_name": None,
            "company_name": None,
            "css_url": None
        }},
        multi=True
    )
    print "##### After Deleting bill_plan #####"
    pprint(items)

def delete_bill_plan_mailers(uid):
    items = prod.bill_plan_mailers.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting bill_plan_mailers #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    prod.bill_plan_mailers.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "username": None,
            "login_url": None,
            "oauth_data": None,
            "dc": None,
            "client_id": "deleted",
            "client_secret": "deleted",
            "password": None,
            "access_token": None,
            "list_association": [],
            "login": {}
        }},
        multi=True
    )
    print "##### After Deleting bill_plan_mailers #####"
    pprint(items)

def delete_bill_subscribe(uid):
    items = prod.bill_subscribe.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting bill_subscribe #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    prod.bill_subscribe.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            #"deactivated_on": now(),
            "title": "Deleted",
            "deleted": True
        }},
        multi=True
    )
    print "##### After Deleting bill_subscribe #####"
    pprint(items)

def delete_cart(uid):
    items = prod.cart.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting cart #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    prod.cart.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "description": "Deleted",
            "title": "Deleted"
        }},
        multi=True
    )
    print "##### After Deleting cart #####"
    pprint(items)

def delete_customurl(uid):
    items = prod.customurl.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting customurl #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    prod.customurl.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "url": "Deleted",
            "deleted": True
        }},
        multi=True
    )
    print "##### After Deleting customurl #####"
    pprint(items)

def delete_payment_account(uid):
    items = prod.payment_account.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting payment_account #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    prod.payment_account.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "description": "Deleted",
            "authorization_code": "deleted",
            "nickname": "deleted",
            "method": "deleted",
            "cim_id": None,
            "stripe_id": "deleted",
            "deleted": True
        }},
        multi=True
    )
    print "##### After Deleting payment_account #####"
    pprint(items)

def delete_payout_account(uid):
    items = prod.payout_account.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting payout_account #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    prod.payout_account.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
            "description": "deleted",
            "nickname": "deleted",
            "city": "deleted",
            "payee_name": "deleted",
            "_name": "deleted",
            "state": "deleted",
            "address_two": "deleted",
            "address_one": "deleted",
            "country": "deleted",
            "method": "deleted",
            "zip_code": "deleted",
            "re_email": "deleted"
        }},
        multi=True
    )
    print "##### After Deleting payout_account #####"
    pprint(items)

def delete_plan_history(uid):
    items = prod.plan_history.find({"author_id": uid})
    items = [i for i in items]
    print "##### Before Deleting payout_history #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    prod.plan_history.update(
        {"author_id": uid},
        {"$set": {
            "author.logo": None,
            "author.display_name": "Deleted",
        }},
        multi=True
    )
    print "##### After Deleting payout_history #####"
    pprint(items)

def delete_user(uid):
    items = prod.user.find({"_id": uid})
    items = [i for i in items]
    print "##### Before Deleting user #####"
    pprint(items)
    _delete_remote_data(items, _filter_remote_urls)

    prod.user.update(
       {"_id": uid},
       {"$set": {
           "author": {},
           "username": None,
           "description": None,
           "firstname": "Deleted",
           "unverified_email": [],
           "lastname": "Deleted",
           "is_active": False,
           "sex": "deleted",
           "timezone": "Deleted",
           "logo": None,
           "original_logo": None,
           "open_access": {},
           "deleted": True,
           "middlename": None,
           "salt": "Deleted",
           "email": "deleted",
           "stripe_id": None
       }}
    )
    print "##### After Deleting user #####"
    pprint(items)


# subpub db
# def delete_feed(uid):
#     items = subpub.feed.find({"author_id": uid})
#     items = [i for i in items]
#    _delete_remote_data(items, _filter_remote_urls)
#
#     subpub.feed.update(
#         {"user_id": uid},
#         {"$set": {
#             "author.logo": None,
#             "author.display_name": "Deleted",
#             "text": "deleted",
#             "title": "deleted",
#             "parent_title": "deleted"
#         }}
#     )
#
# def delete_ticker_activity(uid):
#     items = subpub.ticker_activity.find({"author_id": uid})
#     items = [i for i in items]
#     _delete_remote_data(items, _filter_remote_urls)
#
#     subpub.ticker_activity.update(
#         {"author_id": uid},
#         {"$set": {
#             "author.logo": None,
#             "author.display_name": "Deleted",
#             "text": "deleted",
#             "topic_title": "deleted"
#         }}
#     )


FUNC = {
    delete_asset,
    delete_bundle,
    delete_community,
    delete_coupon,
    delete_feedback,
    delete_jcomment,
    delete_jontent,
    delete_month,
    delete_quizanswer,
    delete_quizquestion,
    delete_quizresult,
    delete_response,
    delete_session,
    delete_signup_code,
    delete_siminar,
    delete_step,
    delete_subscribe,
    delete_testimonial,
    delete_bill_plan,
    delete_bill_plan_mailers,
    delete_bill_subscribe,
    delete_cart,
    delete_customurl,
    delete_payment_account,
    delete_payout_account,
    delete_plan_history,
    delete_user,

    # delete_feed,
    # delete_ticker_activity
}

if __name__ == "__main__":
    user_id = sys.argv[1]
    print user_id

    for fn in FUNC:
fn(user_id)